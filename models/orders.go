package models

type Order struct {
	Product string `json:"product"`
	HashTx  string `json:"hashTx"`
	Price   int    `json:"price"`
	UserId  string `json:"userId"`
}
