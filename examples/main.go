package main

import (
	"fmt"

	"github.com/sociafill/proxify/checker/wtfismyip"
	"github.com/sociafill/proxify/pool"
)

func main() {
	proxiesPool := pool.NewProxyPool(wtfismyip.WtfProxyChecker{})
	fmt.Printf("Proxies pool created: %v\n", proxiesPool)
	proxiesPool.Add("socks5://174.75.238.87:16412")
}
