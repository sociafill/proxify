package common

import "time"

// ProxyCheckResult contains result of proxy check runned once
type ProxyCheckResult struct {
	Delay       time.Duration
	IP          string
	CountryCode string
	Location    string
	Hostname    string
	ISP         string
	TorExit     bool
}
