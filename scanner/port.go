package scanner

import (
// "fmt"
// "time"
// "log"
)

// Port contains info about a port
type Port struct {
	IsOpen      bool   `json:"isopen"`
	Number      int    `json:"number"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

// NewPort constructor
func NewPort(isOpen bool, number int, service string, description string) *Port {
	return &Port{isOpen, number, service, description}
}
