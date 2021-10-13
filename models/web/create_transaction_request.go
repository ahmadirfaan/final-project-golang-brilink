package web

type CreateTransactionRequest struct {
	TransactionTypeId uint   `json:"transactionTypeId" validate:"required"`
	CustomerId        uint   `json:"customerId" validate:"required"`
	AgentId           uint   `json:"agentId" validate:"required"`
	Address           string `json:"address" validate:"required"`
	DistrictId        uint   `json:"districtId" validate:"required,min=7,numeric"`
	Amount            uint64 `json:"amount" validate:"required"`
}
