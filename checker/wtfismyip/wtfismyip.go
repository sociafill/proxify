package wtfismyip

import "github.com/sociafill/proxify/checker/common"

// ProxyCheckResult is main struct to unmarshal service response
type ProxyCheckResult struct {
	common.ProxyCheckResult
	IP          string `json:"YourFuckingIPAddress"`
	Location    string `json:"YourFuckingLocation"`
	Hostname    string `json:"YourFuckingHostname"`
	ISP         string `json:"YourFuckingISP"`
	TorExit     string `json:"YourFuckingTorExit"`
	CountryCode string `json:"YourFuckingCountryCode"`
}
