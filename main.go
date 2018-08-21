package main

import (
	"fmt"

	"github.com/sociafill/proxify/pool"
)

func main() {
	proxyPool := pool.NewProxyPool()
	proxyPool.Add("socks5://202.87.31.203:1080")
	fmt.Printf("Proxies pool created: %v\n", proxyPool)
}
