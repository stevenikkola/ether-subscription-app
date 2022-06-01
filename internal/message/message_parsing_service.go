package message

import (
	"encoding/json"
)

func ParseMessage(notification []byte) TxContents {
	bxMessage := BloxrouteMessage{}
	err := json.Unmarshal(notification, &bxMessage)
	if err != nil {
		panic(err)
	}
	txn := bxMessage.Params.Result.TxContents
	txn.TxHash = bxMessage.Params.Result.TxHash
	return txn
}

type BloxrouteMessage struct {
	Params SubscriptionParams `json:"params"`
}

type SubscriptionParams struct {
	Subscription string             `json:"subscription"`
	Result       SubscriptionResult `json:"result"`
}

type SubscriptionResult struct {
	TxHash     string     `json:"txHash"`
	TxContents TxContents `json:"txContents"`
}

type TxContents struct {
	TxHash               string
	From                 string `json:"from"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	Value                string `json:"value"`
	Input                string `json:"input"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Nonce                string `json:"nonce"`
	R                    string `json:"r"`
	S                    string `json:"s"`
	V                    string `json:"v"`
	To                   string `json:"to"`
}
