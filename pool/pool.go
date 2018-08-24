package pool

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/sociafill/proxify/checker"

	"github.com/sociafill/proxify/checker/wtfismyip"
	"github.com/sociafill/proxify/proxy"
	netProxy "golang.org/x/net/proxy"
)

// ProxyPool is an objects pool for proxies management
type ProxyPool struct {
	checker checker.ProxyChecker
	proxies []proxy.Proxy
}

// NewProxyPool returns instantiated pool
func NewProxyPool(proxyChecker checker.ProxyChecker) ProxyPool {
	proxyPool := ProxyPool{proxies: make([]proxy.Proxy, 0)}
	proxyPool.checker = proxyChecker
	return proxyPool
}

// Add allows add proxies to the pool
func (proxyPool *ProxyPool) Add(URL string) {
	proxyObject := proxy.Proxy{URL: URL}
	proxyPool.proxies = append(proxyPool.proxies, proxyObject)
	proxyPool.runChecker(&proxyObject)
}

func (proxyPool *ProxyPool) runChecker(proxyObject *proxy.Proxy) error {
	httpClient, err := createHTTPClient(proxyObject)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create http client:", err)
		return err
	}

	// @TODO Move to DI
	checker := wtfismyip.WtfProxyChecker{}
	checkResult, err := checker.Check(httpClient)
	if err != nil {
		fmt.Printf("Checking failed: %v\n", err)
		return nil
	}
	fmt.Printf("Check result: %v\n", checkResult)
	return nil
}

func createHTTPClient(proxyObject *proxy.Proxy) (*http.Client, error) {
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

func createDialer(proxyObject *proxy.Proxy) (netProxy.Dialer, error) {
	httpProxyURI, _ := url.Parse(proxyObject.URL)
	dialer, err := netProxy.FromURL(httpProxyURI, netProxy.Direct)
	if err != nil {
		return nil, err
	}
	return dialer, nil
}
