package main

import (
	"fmt"

	"github.com/sociafill/proxify/pool"
)

func main() {
	proxiesPool := pool.NewProxyPool()
	fmt.Printf("Proxies pool created: %v\n", proxiesPool)
	proxiesPool.Add("socks5://101.0.76.120:8080")
	proxiesPool.Add("socks5://168.232.198.81:6667")
}
