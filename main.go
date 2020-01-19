package main

import (
	"runtime"
	"time"
	"github.com/viiftw/glance/scanner"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	start := time.Now()

	s := scanner.NewScanner("localhost", 500*time.Millisecond, 30000, "tcp")
	s.Scan(0,65535)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}