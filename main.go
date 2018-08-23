package main

import (
	"fmt"

	"github.com/sociafill/proxify/checker/wtfismyip"
	"github.com/sociafill/proxify/pool"
)

func main() {
	checker := wtfismyip.WtfProxyChecker{}
	proxyPool := pool.NewProxyPool(checker)
	proxyPool.Add("socks5://165.227.130.164:1090")
	fmt.Printf("Proxies pool created: %v\n", proxyPool)
}
