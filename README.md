# orientwalt-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建HTDF.ini文件，编辑如下内容：

```ini

# transaction type
txType = "auth/StdTx"
# message type
msgSend = "htdfservice/send"
msgVote = "htdfservice/vote"
msgDelegate = "htdfservice/delegate"
# message choose 1-send  2-vote  3-delegate
msgType = 1


# mainnet rest api url
mainnetRestAPI = "http://ip:port"
#"http://47.57.26.144:20014"
# mainnet node api url
mainnetNodeAPI = ""
# chain id
mainnetChainID = "mainchain"
# mainnet denom
mainnetDenom = "satoshi"

# testnet rest api url
testnetRestAPI = ""
# testnet node api url
testnetNodeAPI = ""
# chain id
testnetChainID = ""
# testnet denom
testnetDenom = ""

# Is network test?
isTestNet = false

# scan mempool or not
isScanMemPool = false

# fixed fee to pay in satoshi
fixedFee = 3000000

gasLimit = 200000
gasPrice = 100

# Cache data file directory, default = "", current directory: ./data
dataDir = ""
```
