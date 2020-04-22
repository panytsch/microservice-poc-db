package db

import (
	"log"
)

const CreateTransaction string = "createTransaction"

type CreateTransactionProcedure struct {
	StoredProcedure
	Config *CreateTransactionProcedureConfig
}

type CreateTransactionProcedureConfig struct {
	SPConfig
}

func NewCreateTransactionProcedure() *CreateTransactionProcedure {
	res := &CreateTransactionProcedure{}
	res.DB = DB
	config := &CreateTransactionProcedureConfig{}
	config.Name = CreateTransaction
	res.Config = config
	return res
}

type CreateTransactionProcedureResult struct {
	Result     *Transaction
	ReturnInfo CreateTransactionProcedureReturnInfo
}

type CreateTransactionProcedureReturnInfo struct {
	SPReturnInfo
}

func (sp *CreateTransactionProcedure) Run(userID uint, status TransactionStatus, amount TransactionAmount) *CreateTransactionProcedureResult {
	result := new(CreateTransactionProcedureResult)
	result.Result = new(Transaction)
	result.ReturnInfo = CreateTransactionProcedureReturnInfo{}
	sql := "exec " + sp.Config.Name + " ?,?,?"
	rows, err := sp.DB.Raw(sql, userID, status, int(amount)).Rows()
	if err != nil {
		result.ReturnInfo.ReturnCode = ErrorWhileRun
		log.Printf("While run %v\n", err)
		return result
	}
	for rows.Next() {
		err = sp.DB.ScanRows(rows, result.Result)
		if err != nil {
			result.ReturnInfo.ReturnCode = ErrorWhileRowScan
			log.Printf("While scan %v\n", err)
			return result
		}
	}
	if !rows.NextResultSet() {
		result.ReturnInfo.ReturnCode = NoReturnDataReceived
		return result
	}
	rows.Next()
	err = rows.Scan(&result.ReturnInfo.ReturnCode)
	if err != nil {
		log.Printf("While scan info %v\n", err)
		result.ReturnInfo.ReturnCode = NoReturnInfoReceived
	}

	return result
}

func (result *CreateTransactionProcedureResult) IsSuccess() bool {
	return result.ReturnInfo.ReturnCode == 1
}
