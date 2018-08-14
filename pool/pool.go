package pool

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
func (s *ProxyPool) Add(URL string) {
	proxy := Proxy{URL: URL}
	s.proxies = append(s.proxies, proxy)
}
