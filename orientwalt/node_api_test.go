package orientwalt

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)
const nodeUrl = "http://ip:port"

func Test_getBlockHeight(t *testing.T) {
	c := NewClient(nodeUrl, false)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {
	c := NewClient(nodeUrl, false)

	height := uint64(155960)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
func Test_tmp(t *testing.T) {
	test, err := time.Parse(time.RFC3339Nano, "2019-05-08T02:13:41.937681458Z")
	fmt.Println(err)
	fmt.Println(test.Unix())
}
func Test_getBalance(t *testing.T) {
	c := NewClient(nodeUrl, true)

	address := "htdf1s3hta89j4ykwsaad5sg2ag7twg6lfxqj5mwgup" // 80000000
	//"htdf13nqfr68hclx92lm066lhwxyafpfak3ydar9rrf" // 80000000
	//htdf1s3hta89j4ykwsaad5sg2ag7twg6lfxqj5mwgup       // 93000000
	r, err := c.getBalance(address, "satoshi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {
	c := NewClient(nodeUrl, true)
	txid := "85ef5ff52d2a034e7b3a6ff517ef2553442013274a87bc86e8dd8e554c75545d"
txid = strings.ToUpper(txid)
	path := "/txs/" + txid
	r, err := c.Call(path, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	trx := NewTransaction(r, "auth/StdTx", "htdfservice/send", "satoshi")

	fmt.Println(trx)
}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	c := NewClient("http://localhost:9922/", false)
	result, err := c.getMultiAddrTransactions("auth/StdTx", "cosmos-sdk/MsgSend", "uatom", 0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_getBlockByHeight(t *testing.T) {
	height := uint64(2002)
	c := NewClient(nodeUrl, false)
	result, err := c.getBlockByHeight(height)
	if err != nil {
		t.Error("get block failed!")
	} else {
		fmt.Println(result)
	}
}

func Test_sequence(t *testing.T) {
	addr := "htdf135fj6kmw4el87el6qtl7gwjc9pqqn4hhv35scw"
	c := NewClient(nodeUrl, true)
	accountnumber, sequence, err := c.getAccountNumberAndSequence(addr)
	fmt.Println(err)
	fmt.Println(accountnumber)
	fmt.Println(sequence)
}
