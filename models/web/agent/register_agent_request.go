package web

type RegisterAgentRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
	AgentName   string `json:"agentName" validate:"required,min=5"`
	NoHandphone string `json:"noHandphone" validate:"required,min=9,max=12,numeric"`
	District    string `json:"district" validate:"required"`
	City        string `json:"city" validate:"required"`
	Province    string `json:"province" validate:"required"`
	Address     string `json:"address" validate:"required"`
}
