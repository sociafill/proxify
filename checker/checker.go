package checker

import (
	"net/http"
	"time"
)

// ProxyCheckResult contains result of proxy check runned once
type ProxyCheckResult struct {
	Delay       time.Duration
	IP          string
	CountryCode string
	Location    string
	Hostname    string
	ISP         string
	IsTorExit   bool
}

// ProxyChecker is an interface for checkers
type ProxyChecker interface {
	Check(*http.Client) (ProxyCheckResult, error)
}
