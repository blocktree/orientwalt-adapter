package orientwaltTransaction

import (
	"encoding/base64"
)

type Pub struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewPub(pubkey []byte, curveType string) Pub {
	return Pub{
		Type:  curveType,
		Value: base64.StdEncoding.EncodeToString(pubkey),
	}
}

type Sig struct {
	Pubkey        Pub    `json:"pub_key"`
	Signature     string `json:"signature"`
}
type Sigs []Sig

func NewSig(signature []byte, pubkey Pub) Sig {
	return Sig{
		Pubkey:        pubkey,
		Signature:     base64.StdEncoding.EncodeToString(signature),
	}
}


type AmountSend struct {
	Denom string `json:"denom"`
	Amount string `json:"amount"`
}

type AmountSends []AmountSend

func NewAmountSend(amount, denom string) AmountSend {
	return AmountSend{
		Denom:  denom,
		Amount: amount,
	}
}

type FeeSend struct {
	GasWanted string `json:"gas_wanted"`
	GasPrice string `json:"gas_price"`
}


type MsgValue struct {
	From string `json:"From"`
	To string `json:"To"`
	Amount AmountSends `json:"Amount"`
	Data string `json:"Data"`
	GasPrice string `json:"GasPrice"`
	GasWanted string `json:"GasWanted"`
}


func NewMsgValue(from, to, denom, amount, gasPrice, gasWanted string) MsgValue {
	return MsgValue{
		From:      from,
		To:        to,
		Amount:    AmountSends{NewAmountSend(amount, denom)},
		Data:      "",
		GasPrice:  gasPrice,
		GasWanted: gasWanted,
	}
}

type Msg struct {
	Type string `json:"type"`
	MsgValue MsgValue `json:"value"`
}
type Msgs []Msg

func NewMsg(msgType, from, to, denom, amount, gasPrice, gasWanted string) Msg {
	return Msg{
		Type:     msgType,
		MsgValue:NewMsgValue(from, to, denom, amount, gasPrice, gasWanted),
	}
}


type JsonValue struct {
	Msgs Msgs `json:"msg"`
	Fee FeeSend `json:"fee"`
	Sigs Sigs `json:"signatures"`
	Memo string `json:"memo"`
}


func NewJsonValue(msgType,pubType, memo,  from, to, denom, amount, gasPrice, gasWanted string, pub, sig []byte) JsonValue {
	return JsonValue{
		Msgs: Msgs{NewMsg(msgType, from, to, denom, amount, gasPrice, gasWanted)},
		Fee:  FeeSend{
			GasWanted: gasWanted,
			GasPrice:  gasPrice,
		},
		Sigs: Sigs{NewSig(sig, NewPub(pub, pubType))},
		Memo: memo,
	}
}


type JsonSend struct {
	Type string `json:"type"`
	Value JsonValue `json:"value"`
}


func NewJsonSend(txType, msgType,pubType,memo, from, to, denom , amount, gasPrice, gasWanted string, pub, sig []byte) JsonSend {
	return JsonSend{
		Type:  txType,
		Value: NewJsonValue(msgType, pubType, memo, from,to, denom, amount, gasPrice, gasWanted, pub, sig),
	}
}
