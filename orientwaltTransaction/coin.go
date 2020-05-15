package orientwaltTransaction

import (
	"strconv"
)

type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}
type Coins []Coin

func NewCoin(denom string, amount int64) Coin {
	return Coin{
		Denom:  denom,
		Amount: strconv.FormatInt(amount, 10),
	}
}
