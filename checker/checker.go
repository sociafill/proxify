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

// // ProxyChecker is first-class function for one more abstraction
// type ProxyChecker = func(httpClient *http.Client) (*common.ProxyCheckResult, error)

// // WtfProxyChecker uses https://wtfismyip.com/ API
// var WtfProxyChecker = func(httpClient *http.Client) (common.ProxyCheckResult, error) {
// 	result := common.ProxyCheckResult{}
// 	req, err := http.NewRequest("GET", "https://wtfismyip.com/json", nil)
// 	if err != nil {
// 		log.Printf("can't create request: %s\n", err)
// 		return result, err
// 	}

// 	start := time.Now()
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		log.Printf("can't GET page: %s\n", err)
// 		return result, err
// 	}
// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("error reading body: %s\n", err)
// 		return result, err
// 	}

// 	var proxyCheckResult wtfismyip.ProxyCheckResult
// 	json.Unmarshal(b, &proxyCheckResult)
// 	log.Printf("Received data %v\n", proxyCheckResult)
// 	proxyCheckResult.Delay = time.Since(start)

// 	return result, nil
// }
