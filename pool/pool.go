package pool

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/sociafill/proxify/checker"
	"golang.org/x/net/proxy"
)

// Proxy contains information required for golang.org/x/net/proxy instantiation
type Proxy struct {
	URL string
}

// ProxyPool is an objects pool for proxies management
type ProxyPool struct {
	proxies []Proxy
}

// NewProxyPool returns instantiated pool
func NewProxyPool() ProxyPool {
	proxyPool := ProxyPool{proxies: make([]Proxy, 0)}
	return proxyPool
}

// Add allows add proxies to the pool
func (proxyPool *ProxyPool) Add(URL string) {
	proxyObject := Proxy{URL: URL}
	proxyPool.proxies = append(proxyPool.proxies, proxyObject)
	proxyPool.runChecker(&proxyObject)
}

func (proxyPool *ProxyPool) runChecker(proxyObject *Proxy) error {
	httpClient, err := createHTTPClient(proxyObject)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create http client:", err)
		return err
	}

	checkResult, err := checker.WtfProxyChecker(httpClient)
	fmt.Printf("Check result: %v\n", checkResult.Delay)
	return nil
}

func createHTTPClient(proxyObject *Proxy) (*http.Client, error) {
	dialer, err := createDialer(proxyObject)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		return nil, err
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial
	httpClient := &http.Client{Transport: httpTransport}
	return httpClient, nil
}

func createDialer(proxyObject *Proxy) (proxy.Dialer, error) {
	httpProxyURI, _ := url.Parse(proxyObject.URL)
	dialer, err := proxy.FromURL(httpProxyURI, proxy.Direct)
	if err != nil {
		return nil, err
	}
	return dialer, nil
}
