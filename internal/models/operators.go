package models

type GetOperatorsModel struct {
	OperatorID   string `json:"operator_id"`
	OperatorName string `json:"operator_name"`
}

type CreateOperatorModel struct {
	OperatorName string `json:"operator_name"`
}
