package db

import (
	"log"
)

const CreatePayment string = "createPayment"

type CreatePaymentProcedure struct {
	StoredProcedure
	Config *CreatePaymentProcedureConfig
}

type CreatePaymentProcedureConfig struct {
	SPConfig
}

func NewCreatePaymentProcedure() *CreatePaymentProcedure {
	res := &CreatePaymentProcedure{}
	res.DB = DB
	config := &CreatePaymentProcedureConfig{}
	config.Name = CreatePayment
	res.Config = config
	return res
}

type CreatePaymentProcedureResult struct {
	Result     *Payment
	ReturnInfo CreatePaymentProcedureReturnInfo
}

type CreatePaymentProcedureReturnInfo struct {
	SPReturnInfo
}

func (sp *CreatePaymentProcedure) Run(userID uint, status PaymentStatus, amount PaymentAmount) *CreatePaymentProcedureResult {
	result := new(CreatePaymentProcedureResult)
	result.Result = new(Payment)
	result.ReturnInfo = CreatePaymentProcedureReturnInfo{}
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

func (result *CreatePaymentProcedureResult) IsSuccess() bool {
	return result.ReturnInfo.ReturnCode == 1
}
