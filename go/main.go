package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"github.com/panytsch/microservice-poc-db/go/rest/server"
	"sync"
)

func main() {
	core.ConnectDB()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go server.RunRestServer(wg)

	wg.Wait()
	//sp := test.NewTwoDataSetsProcedure(core.DB)
	//res := sp.Run()
	//log.Printf("sp result %v", res)
}
