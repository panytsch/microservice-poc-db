package main

import (
	"github.com/panytsch/microservice-poc-db/go/rest_v1"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go rest_v1.RunRestServer(wg)

	wg.Wait()
}
