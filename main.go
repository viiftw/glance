package main

import (
	"net/http"
	"runtime"
	"time"
	"github.com/viiftw/glance/scanner"
	"fmt"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	valid "github.com/asaskevich/govalidator"
)

// Error contains information about an API error.
type Error struct {
	Code    int `json:"code"`
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
	if !validateInput(target) {
		// c.AbortWithStatus(400)
		err := &Error{
			Code: 400,
			Message: "Invalid hostname or ip",
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("Scanning ", target)

	start := time.Now()

	s := scanner.NewScanner(target, 500*time.Millisecond, 30000, "tcp")
	s.Scan(1, 65535)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	s.Result.UpdateTimeComplete(elapsed.Seconds())

	c.JSON(200, s.Result)
}

func validateInput(target string) bool {
	if valid.IsDNSName(target) || valid.IsIPv4(target) || valid.IsIPv6(target) {
		if scanner.GetIP(target) == scanner.UNKNOWN {
			return  false
		}
		return true
	}
	return  false
}