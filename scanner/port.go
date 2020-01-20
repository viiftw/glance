package scanner

import (
	// "fmt"
	// "time"
	// "log"
)

// Port contains info about a port
type Port struct {
	IsOpen	bool
	Number	int
	Service	string
	Description string
}

// NewPort constructor
func NewPort(isOpen bool, number int, service string, description string) *Port {
	return &Port{isOpen, number, service, description}
}