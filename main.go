package main

import (
	"runtime"
	"time"
	"github.com/viiftw/glance/scanner"
	"fmt"
	// "encoding/json"
	"github.com/gin-gonic/gin"
)

// Error contains information about an API error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	r := gin.Default()
	r.GET("/scan/:host", scanHandler)
	r.Run(":8686")


}

func scanHandler(c *gin.Context) {
	target := c.Param("host")
	fmt.Println("Scanning ", target)

	start := time.Now()

	s := scanner.NewScanner(target, 500*time.Millisecond, 30000, "tcp")
	s.Scan(1, 65535)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	// timeComplete, _ := fmt.Printf("%s", elapsed)
	s.Result.UpdateTimeComplete(elapsed.Seconds())

	c.JSON(200, s.Result)
}