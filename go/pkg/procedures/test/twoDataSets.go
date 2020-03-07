package test

import (
	"github.com/jinzhu/gorm"
	"github.com/panytsch/go_poc/mssql/go/pkg/core"
	"log"
)

const Name string = "twoDataSets"

type TwoDataSetsProcedure struct {
	core.StoredProcedure
	Config *TwoDataSetsProcedureConfig
}

type TwoDataSetsProcedureConfig struct {
	core.SPConfig
}

func NewTwoDataSetsProcedure(db *gorm.DB) *TwoDataSetsProcedure {
	res := &TwoDataSetsProcedure{}
	res.DB = db
	config := &TwoDataSetsProcedureConfig{}
	config.Name = Name
	res.Config = config
	return res
}

type TwoDataSetsProcedureResult struct {
	Result     *TwoDataSetsProcedureReturnData
	ReturnInfo *TwoDataSetsProcedureReturnInfo
}

type TwoDataSetsProcedureReturnInfo struct {
	core.SPReturnInfo
}

type TwoDataSetsProcedureReturnData struct {
	One   int
	Two   string
	Three int
}

func (sp *TwoDataSetsProcedure) Run() *TwoDataSetsProcedureResult {
	result := initResult()
	sql := "exec " + sp.Config.Name
	rows, err := sp.DB.Raw(sql).Rows()
	if err != nil {
		result.ReturnInfo.ReturnCode = core.ErrorWhileRun
		log.Printf("While run %v\n", err)
		return result
	}
	for rows.Next() {
		err = sp.DB.ScanRows(rows, result.Result)
		if err != nil {
			result.ReturnInfo.ReturnCode = core.ErrorWhileRowScan
			log.Printf("While scan %v\n", err)
			return result
		}
	}
	if !rows.NextResultSet() {
		result.ReturnInfo.ReturnCode = core.NoReturnDataReceived
		return result
	}
	rows.Next()
	err = rows.Scan(&result.ReturnInfo.ReturnCode)
	if err != nil {
		log.Printf("While scan info %v\n", err)
		result.ReturnInfo.ReturnCode = core.NoReturnInfoReceived
	}

	return result
}

func initResult() *TwoDataSetsProcedureResult {
	result := &TwoDataSetsProcedureResult{}
	result.ReturnInfo = &TwoDataSetsProcedureReturnInfo{}
	result.Result = &TwoDataSetsProcedureReturnData{}
	return result
}
