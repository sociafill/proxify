package main

import (
	"fmt"

	"github.com/sociafill/proxify/pool"
)

func main() {
	proxyPool := pool.NewProxyPool()
	proxyPool.Add("socks5h://111.231.88.18:1080")
	fmt.Printf("Proxies pool created: %v\n", proxyPool)
}
