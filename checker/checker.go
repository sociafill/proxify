package checker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// ProxyCheckResult contains result of proxy check runned once
type ProxyCheckResult struct {
	Delay time.Duration
}

// ProxyChecker is first-class function for one more abstraction
type ProxyChecker = func(httpClient *http.Client) (*ProxyCheckResult, error)

// WtfProxyChecker uses https://wtfismyip.com/ API
var WtfProxyChecker = func(httpClient *http.Client) (ProxyCheckResult, error) {
	result := ProxyCheckResult{}
	req, err := http.NewRequest("GET", "https://wtfismyip.com/json", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		return result, err
	}

	start := time.Now()
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		return result, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		return result, err
	}
	result.Delay = time.Since(start)
	fmt.Println(string(b))
	return result, nil
}
