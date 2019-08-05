package models

type TxId struct {
	Result bool   `json:"result" xml:"result"`
	TxId   string `json:"txId" xml:"txId"`
}
