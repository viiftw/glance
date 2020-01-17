package main

import (
	"time"
	"github.com/viiftw/glance/scanner"
)

func main() {
	s := scanner.NewScanner("google.com", 500*time.Millisecond, 30000, "tcp")
	s.Scan(0,65535)
}