package main

type Ping struct {
	Version string `json:"version"`
	Health  string `json:"health"`
	Msg     string `json:"msg"`
}

type Bank struct {
	OwningUserId  string `json:"owningUserId"`
	BankId        string `json:"bankId"`
	BankName      string `json:"bankName"`
	AccountNumber string `json:"accountNumber"`
}
