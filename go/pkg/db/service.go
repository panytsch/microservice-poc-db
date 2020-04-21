package db

import (
	"github.com/jinzhu/gorm"
)

type ReturnCode int

const (
	ErrorWhileRun     ReturnCode = iota + 5000 // start from 5000
	ErrorWhileRowScan                          //5001 ...
	NoReturnDataReceived
	NoReturnInfoReceived
)

type StoredProcedure struct {
	DB     *gorm.DB
	Config *SPConfig
}

type SPConfig struct {
	Name string
}

type SPReturnInfo struct {
	ReturnCode    ReturnCode
	ReturnMessage string
}
