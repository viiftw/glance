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
}

// Host contains the scan results and information about a host.
type Host struct {
	Addr  string
	IP string
	IsUp  bool
	Ports []*Port
	Vendor	string
	OSInfo	string
	Mac	string
	TimeComplete time.Duration
}

// Service contains the info about service
type Service struct {
	Name string
	Version string
	Description string
}

// Port contains info about a port
type Port struct {
	IsOpen	bool
	Number	int
	ServiceInfo	Service
}

// Error contains information about an API error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewScanner constructor
func NewScanner(host string, timeout time.Duration, concurrent int64, protocol string) *Scanner {
	protocol = "tcp"
	return &Scanner{host, timeout, semaphore.NewWeighted(concurrent), protocol}
}

func (s *Scanner) SetConcurrent(concurrent int64) {
	s.Concurrent = semaphore.NewWeighted(concurrent)
}

func (s *Scanner) SetTimeout(timeout time.Duration) {
	s.Timeout = timeout
}

func (s *Scanner) SetProtocol(protocol string) {
	s.Protocol = protocol
}

func scanTCP(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", target)
	if err != nil {
		log.Println(err)
		return
	}
	
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), timeout)

	if err != nil {
		// dial tcp <host:port>: connectex: No connection could be made because the target machine actively refused it.
		// dial tcp <host:port>: connectex: The requested address is not valid in its context.
		// dial tcp <host:port>: i/o timeout
		// dial tcp <host:port>: bind: An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full.

		if strings.Contains(err.Error(), "bind: An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full.") {		
			time.Sleep(timeout)
			scanTCP(ip, port, timeout)
		} else {
			// fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (s *Scanner) Scan(startPort int, endPort int) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := startPort; port <= endPort; port++ {
		s.Concurrent.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer s.Concurrent.Release(1)
			defer wg.Done()
			scanTCP(s.Host, port, s.Timeout)
		}(port)
	}
}

func (s Scanner) renderAddr(port int) string {
	return fmt.Sprintf("%s:%d", s.Host, port)
}