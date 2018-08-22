package main

import (
	"fmt"

	"github.com/sociafill/proxify/pool"
)

func main() {
	proxyPool := pool.NewProxyPool()
	proxyPool.Add("socks5://165.227.130.164:1090")
	fmt.Printf("Proxies pool created: %v\n", proxyPool)
}
