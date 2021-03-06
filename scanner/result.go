package scanner

import (
// "time"
)

// Host contains the scan results and information about a host.
type Host struct {
	Addr         string  `json:"addr"`
	IP           string  `json:"ip"`
	IsUp         bool    `json:"isup"`
	Ports        []*Port `json:"ports"`
	Vendor       string  `json:"vendor"`
	OSInfo       string  `json:"osinfo"`
	Mac          string  `json:"mac"`
	TimeComplete float64 `json:"timecomplete"`
}

// NewHost constructor
func NewHost(host string) *Host {
	return &Host{Addr: host}
}

// UpdateStatus update host up or down
func (h *Host) UpdateStatus(isUp bool) {
	h.IsUp = isUp
}

// UpdateIP update ip result
func (h *Host) UpdateIP(ip string) {
	h.IP = ip
}

// UpdateTimeComplete update time scan complete
func (h *Host) UpdateTimeComplete(time float64) {
	h.TimeComplete = time
}

// UpdateInfo update info about result
func (h *Host) UpdateInfo(vendor string, osInfo string, mac string) {
	h.Vendor = vendor
	h.OSInfo = osInfo
	h.Mac = mac
}

// UpdatePort add a port open to result
func (h *Host) UpdatePort(port *Port) {
	h.Ports = append(h.Ports, port)
}
