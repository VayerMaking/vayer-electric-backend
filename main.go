package main

import (
	"fmt"
	"vayer-electric-backend/env"
)

func main() {
    fmt.Println(env.REQUEST_TIMEOUT)
	fmt.Println(env.SHUTDOWN_TIMEOUT)
	fmt.Println(env.STATSD_FLUSH)
}