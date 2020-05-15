package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
	dbFilePath     = filepath.Join("data", "db")
	dbFileName     = "blockchain-HTDF.db"
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"HTDF",
	}
	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO HTDF", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
	account := &openwallet.AssetsAccount{Alias: "HELLO HTDF", WalletID: walletID, Required: 1, Symbol: "HTDF", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
	accountID := "6jmmNwXp62dV4KvD9XKHGyWPf2SXbpiUnLo9S5S7Bnky"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
	accountID := "Mbc4KA7R9v98QC2a3aGpmHcbrsCzAjrMtvdiGX6irPe"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
		log.Info("address[", i, "] :", w.PublicKey)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}

//walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
//accountID := "6jmmNwXp62dV4KvD9XKHGyWPf2SXbpiUnLo9S5S7Bnky"
//"htdf13nqfr68hclx92lm066lhwxyafpfak3ydar9rrf" // 80000000
//htdf1s3hta89j4ykwsaad5sg2ag7twg6lfxqj5mwgup   // 93000000


//walletID := "W7nU4HF42MpBJc7xh3dB8nedrzqDjVRfBb"
//accountID := "Mbc4KA7R9v98QC2a3aGpmHcbrsCzAjrMtvdiGX6irPe"
//htdf135fj6kmw4el87el6qtl7gwjc9pqqn4hhv35scw