package wtfismyip

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sociafill/proxify/checker"
)

// WtfProxyChecker is an check based on wtfismyip.com service
type WtfProxyChecker struct {
}

// Check uses API of wtfismyip.com to check http transport of specified client
func (WtfProxyChecker) Check(httpClient *http.Client) (checker.ProxyCheckResult, error) {

	var result checker.ProxyCheckResult

	req, err := http.NewRequest("GET", "https://wtfismyip.com/json", nil)
	if err != nil {
		log.Printf("can't create request: %s\n", err)
		return result, err
	}

	start := time.Now()
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("can't GET page: %s\n", err)
		return result, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading body: %s\n", err)
		return result, err
	}

	var proxyCheckResult wtfProxyCheckResult
	json.Unmarshal(b, &proxyCheckResult)
	result.Delay = time.Since(start)

	log.Printf("Received data %v\n", result)

	return result, nil
}

// ProxyCheckResult is main struct to unmarshal service response
type wtfProxyCheckResult struct {
	IP          string `json:"YourFuckingIPAddress"`
	Location    string `json:"YourFuckingLocation"`
	Hostname    string `json:"YourFuckingHostname"`
	ISP         string `json:"YourFuckingISP"`
	TorExit     string `json:"YourFuckingTorExit"`
	CountryCode string `json:"YourFuckingCountryCode"`
}
