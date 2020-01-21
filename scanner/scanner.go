package scanner

import (
	"context"
	"fmt"
	"net"
	// "os/exec"
	// "strconv"
	"strings"
	"sync"
	// "runtime"
	"time"
	"golang.org/x/sync/semaphore"
	"log"
)

// Scanner is struct of scanner object
type Scanner struct {
	Host string
	Timeout time.Duration
	Concurrent *semaphore.Weighted
	Protocol string
	Result *Host
}

// NewScanner constructor
func NewScanner(host string, timeout time.Duration, concurrent int64, protocol string) *Scanner {
	protocol = "tcp"
	return &Scanner{host, timeout, semaphore.NewWeighted(concurrent), protocol, NewHost(host)}
}

// SetConcurrent set max goroutine for scanner
func (s *Scanner) SetConcurrent(concurrent int64) {
	s.Concurrent = semaphore.NewWeighted(concurrent)
}

// SetTimeout set timeout for net.DialTimeout
func (s *Scanner) SetTimeout(timeout time.Duration) {
	s.Timeout = timeout
}

// SetProtocol set protocol for scanner
func (s *Scanner) SetProtocol(protocol string) {
	s.Protocol = protocol
}

func scanTCP(ip string, port int, timeout time.Duration) int{
	target := fmt.Sprintf("%s:%d", ip, port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", target)
	if err != nil {
		log.Println(err)
		return -1
	}
	
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), timeout)

	if err != nil {
		if strings.Contains(err.Error(), "bind: An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full.") {		
			time.Sleep(timeout)
			scanTCP(ip, port, timeout)
		} else {
			// fmt.Println(port, "closed")
		}
		return -1
	}

	conn.Close()
	// h.IsUp = true
	// fmt.Println(port, "open")
	return port
}

// Scan start scan with startPort and endPort numbers
func (s *Scanner) Scan(startPort int, endPort int) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := startPort; port <= endPort; port++ {
		s.Concurrent.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer s.Concurrent.Release(1)
			defer wg.Done()
			s.handleOpenPort(scanTCP(s.Host, port, s.Timeout))
		}(port)
	}
}

func (s *Scanner) handleOpenPort(portNumber int) {
	if portNumber == -1 {
		return
	}
	service := predictPort(portNumber)
	description := ""
	port := NewPort(true, portNumber, service, description)

	resultsMutex := sync.Mutex{}
	resultsMutex.Lock()
	s.Result.UpdatePort(port)
	if !s.Result.IsUp {
		s.Result.UpdateStatus(true)
		s.Result.UpdateIP(GetIP(s.Host))
	}
	resultsMutex.Unlock()
}

// GetIP return ips of hostname
func GetIP(host string) string {
	ip, err:= net.LookupIP(host)
	if err != nil {
		return UNKNOWN
	}
	return fmt.Sprintf("%s", ip)
}