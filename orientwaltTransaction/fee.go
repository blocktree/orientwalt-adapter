package orientwaltTransaction

import "strconv"

type FeeStruct struct {
	GasPrice string  `json:"gas_price"`
	GasWanted string `json:"gas_wanted"`
}

func NewStdFee(gasPrice, gasWanted int64) FeeStruct {
	return FeeStruct{
		GasPrice:     strconv.FormatInt(gasPrice, 10),
		GasWanted:    strconv.FormatInt(gasWanted, 10),
	}
}
