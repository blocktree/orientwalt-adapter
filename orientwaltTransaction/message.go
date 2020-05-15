package orientwaltTransaction

type MsgSend struct {
	Amount      Coins  `json:"amount"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
}

type Message struct {
	Amount Coins `json:"Amount"`
	Data string `json:"Data"`
	From string `json:"From"`
	GasPrice int64 `json:"GasPrice"`
	GasWanted int64 `json:"GasWanted"`
	To string `json:"To"`
}

func NewMessage(from, to, denom  string, amount, gasPrice, gasWanted int64) Message {
	return Message{
		Amount:    Coins{NewCoin(denom, amount)},
		Data:      "",
		From:      from,
		GasPrice:  gasPrice,//strconv.FormatInt(gasPrice, 10),
		GasWanted: gasWanted,//strconv.FormatInt(gasWanted, 10),
		To:        to,
	}
}
