package gui

// Wallet-related information for the GUI
import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/skycoin/skycoin/src/cipher"
	bip39 "github.com/skycoin/skycoin/src/cipher/go-bip39"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/daemon"
	"github.com/skycoin/skycoin/src/visor"
	"github.com/skycoin/skycoin/src/wallet"

	wh "github.com/skycoin/skycoin/src/util/http" //http,json helpers
)

// WalletRPC wallet rpc
// type WalletRPC struct {
// 	sync.RWMutex
// 	wallets         wallet.Wallets
// 	WalletDirectory string
// 	options         []wallet.Option
// 	firstAddrIDMap  map[string]string // key: first address in wallet, value: wallet id
// }

// NotesRPC note rpc
type NotesRPC struct {
	Notes           wallet.Notes
	WalletDirectory string
}

// Wg use a global for now
// var Wg *WalletRPC

// Ng global note
var Ng *NotesRPC

// InitWalletRPC init wallet rpc
func InitWalletRPC(walletDir string, options ...wallet.Option) {
	// Wg = NewWalletRPC(walletDir, options...)
	Ng = NewNotesRPC(walletDir)
}

// NewNotesRPC new notes rpc
func NewNotesRPC(walletDir string) *NotesRPC {
	rpc := &NotesRPC{}
	if err := os.MkdirAll(walletDir, os.FileMode(0700)); err != nil {
		logger.Panicf("Failed to create notes directory %s: %v", walletDir, err)
	}
	rpc.WalletDirectory = walletDir
	w, err := wallet.LoadNotes(rpc.WalletDirectory)
	if err != nil {
		logger.Panicf("Failed to load all notes: %v", err)
	}
	wallet.CreateNoteFileIfNotExist(walletDir)
	rpc.Notes = w
	return rpc
}

// GetNotesReadable returns readable notes
func (nt *NotesRPC) GetNotesReadable() wallet.ReadableNotes {
	return nt.Notes.ToReadable()
}

// SpendResult represents the result of spending
type SpendResult struct {
	Balance     *wallet.BalancePair        `json:"balance,omitempty"`
	Transaction *visor.ReadableTransaction `json:"txn,omitempty"`
	Error       string                     `json:"error,omitempty"`
}

// Spend TODO
// - split send into
// -- get addresses
// -- get unspent outputs
// -- construct transaction
// -- sign transaction
// -- inject transaction
func Spend(gateway *daemon.Gateway,
	walletID string,
	amt wallet.Balance,
	dest cipher.Address) *SpendResult {
	var tx *coin.Transaction
	var b wallet.BalancePair
	var err error
	for {
		tx, err = gateway.Spend(walletID, amt, dest)
		if err != nil {
			break
		}

		logger.Info("Spend: \ntx= \n %s \n", visor.TransactionToJSON(*tx))

		b, err = gateway.GetWalletBalance(walletID)
		if err != nil {
			err = fmt.Errorf("Get wallet balance failed: %v", err)
			break
		}

		break
	}

	if err != nil {
		return &SpendResult{
			Error: err.Error(),
		}
	}

	rdtx := visor.NewReadableTransaction(&visor.Transaction{Txn: *tx})

	return &SpendResult{
		Balance:     &b,
		Transaction: &rdtx,
	}
}

// Returns the wallet's balance, both confirmed and predicted.  The predicted
// balance is the confirmed balance minus the pending spends.
func walletBalanceHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			wh.Error405(w)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			wh.Error400(w, "id is required")
			return
		}

		b, err := gateway.GetWalletBalance(id)
		if err != nil {
			logger.Error("Get wallet balance failed: %v", err)
			return
		}
		wh.SendOr404(w, b)
	}
}

func getBalanceHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			addrsParam := r.URL.Query().Get("addrs")
			addrsStr := strings.Split(addrsParam, ",")
			addrs := make([]cipher.Address, 0, len(addrsStr))
			for _, addr := range addrsStr {
				a, err := cipher.DecodeBase58Address(addr)
				if err != nil {
					wh.Error400(w, fmt.Sprintf("address %s is invalid: %v", addr, err))
					return
				}
				addrs = append(addrs, a)
			}

			bal, err := gateway.GetAddressesBalance(addrs)
			if err != nil {
				logger.Error("Get balance failed: %v", err)
				wh.Error500(w)
				return
			}

			wh.SendOr404(w, bal)
		}
	}
}

// Creates and broadcasts a transaction sending money from one of our wallets
// to destination address.
func walletSpendHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("id") == "" {
			wh.Error400(w, "Missing wallet_id")
			return
		}

		walletID := r.FormValue("id")
		if walletID == "" {
			wh.Error400(w, "Invalid Wallet Id")
			return
		}
		sdst := r.FormValue("dst")
		if sdst == "" {
			wh.Error400(w, "Missing destination address \"dst\"")
			return
		}
		dst, err := cipher.DecodeBase58Address(sdst)
		if err != nil {
			wh.Error400(w, fmt.Sprintf("Invalid destination address: %v", err))
			return
		}

		scoins := r.FormValue("coins")
		coins, err := strconv.ParseUint(scoins, 10, 64)
		if err != nil {
			wh.Error400(w, "Invalid \"coins\" value")
			return
		}

		var hours uint64
		//MOVE THIS INTO HERE
		ret := Spend(gateway, walletID, wallet.NewBalance(coins, hours), dst)
		if ret.Error != "" {
			logger.Error(ret.Error)
		}

		wh.SendOr404(w, ret)
	}
}

// Create a wallet Name is set by creation date
func notesCreate(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("API request made to create a note")
		note := r.FormValue("note")
		transactionID := r.FormValue("transaction_id")
		newNote := wallet.Note{
			TxID:  transactionID,
			Value: note,
		}
		Ng.Notes.SaveNote(Ng.WalletDirectory, newNote)
		rlt := Ng.GetNotesReadable()
		wh.SendOr500(w, rlt)
	}
}

// Create a wallet Name is set by creation date
func walletCreate(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("API request made to create a wallet")
		seed := r.FormValue("seed")
		label := r.FormValue("label")
		wltName := wallet.NewWalletFilename()
		var wlt wallet.Wallet
		var err error
		// the wallet name may dup, rename it till no conflict.
		for {
			wlt, err = gateway.NewWallet(wltName, wallet.OptSeed(seed), wallet.OptLabel(label))
			if err != nil {
				if strings.Contains(err.Error(), "renaming") {
					wltName = wallet.NewWalletFilename()
					continue
				}

				wh.Error400(w, err.Error())
				return
			}
			break
		}

		rlt := wallet.NewReadableWallet(wlt)
		wh.SendOr500(w, rlt)
	}
}

// method: POST
// url: /wallet/newAddress
// params:
// 		id: wallet id
// 	   num: number of address need to create, if not set the default value is 1
func walletNewAddresses(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			wh.Error405(w)
			return
		}

		wltID := r.FormValue("id")
		if wltID == "" {
			wh.Error400(w, "wallet id is empty")
			return
		}

		// the number of address that need to create, default is 1
		n := 1
		var err error
		num := r.FormValue("num")
		if num != "" {
			n, err = strconv.Atoi(num)
			if err != nil {
				wh.Error400(w, "invalid num value")
				return
			}
		}

		addrs, err := gateway.NewAddresses(wltID, n)
		if err != nil {
			wh.Error400(w, err.Error())
			return
		}

		var rlt = struct {
			Address []string `json:"addresses"`
		}{}

		for _, a := range addrs {
			rlt.Address = append(rlt.Address, a.String())
		}

		wh.SendOr404(w, rlt)
		return
	}
}

// Update wallet label
func walletUpdateHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Update wallet
		id := r.FormValue("id")
		if id == "" {
			wh.Error400(w, "wallet id is empty")
			return
		}

		label := r.FormValue("label")
		if label == "" {
			wh.Error400(w, "label is empty")
			return
		}

		if err := gateway.UpdateWalletLabel(id, label); err != nil {
			wh.Error400(w, fmt.Sprintf("update wallet label failed: %v", err))
			return
		}

		wh.SendOr404(w, "success")
	}
}

// Returns a wallet by id
func walletGet(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			wh.Error405(w)
			return
		}

		id := r.FormValue("id")
		if id == "" {
			wh.Error400(w, fmt.Sprintf("wallet id is empty"))
			return
		}

		wlt, ok := gateway.GetWallet(id)
		if !ok {
			wh.Error400(w, fmt.Sprintf("wallet %s doesn't exist", id))
			return
		}

		wh.SendOr404(w, wlt)
	}
}

// Returns a wallet by ID if GET.  Creates or updates a wallet if POST.
func notesHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ret := wallet.Wallets.ToPublicReadable()
		ret := Ng.GetNotesReadable()
		wh.SendOr404(w, ret)
	}
}

// Returns JSON of unconfirmed transactions for user's wallet
func walletTransactionsHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			wh.Error405(w)
			return
		}

		id := r.FormValue("id")
		if id == "" {
			wh.Error400(w, "wallet id is empty")
			return
		}

		txns, err := gateway.GetWalletUnconfirmedTxns(r.FormValue("id"))
		if err != nil {
			wh.Error400(w, fmt.Sprintf("get wallet unconfirmed transactions failed: %v", err))
			return
		}

		wh.SendOr404(w, txns)
	}
}

// Returns all loaded wallets
func walletsHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ret := wallet.Wallets.ToPublicReadable()
		wlts := gateway.GetWalletsReadable()
		wh.SendOr404(w, wlts)
	}
}

// Loads/unloads wallets from the wallet directory
func walletsReloadHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := gateway.ReloadWallets(); err != nil {
			logger.Error("reload wallet failed: %v", err)
			wh.Error500(w)
			return
		}

		wh.SendOr404(w, "success")
	}
}

// WalletFolder struct
type WalletFolder struct {
	Address string `json:"address"`
}

// Loads/unloads wallets from the wallet directory
func getWalletFolder(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ret := WalletFolder{
			Address: gateway.GetWalletDir(),
		}
		wh.SendOr404(w, ret)
	}
}

// getOutputsHandler get utxos base on the filters in url params.
// mode: GET
// url: /outputs?addrs=[:addrs]&hashes=[:hashes]
// if addrs and hashes are not specificed, return all unspent outputs.
// if both addrs and hashes are specificed, then both those filters are need to be matched.
// if only specify one filter, then return outputs match the filter.
func getOutputsHandler(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var addrs []string
			var hashes []string

			trimSpace := func(vs []string) []string {
				for i := range vs {
					vs[i] = strings.TrimSpace(vs[i])
				}
				return vs
			}

			addrStr := r.FormValue("addrs")
			if addrStr != "" {
				addrs = trimSpace(strings.Split(addrStr, ","))
			}

			hashStr := r.FormValue("hashes")
			if hashStr != "" {
				hashes = trimSpace(strings.Split(hashStr, ","))
			}

			filters := []daemon.OutputsFilter{}
			if len(addrs) > 0 {
				filters = append(filters, daemon.FbyAddresses(addrs))
			}

			if len(hashes) > 0 {
				filters = append(filters, daemon.FbyHashes(hashes))
			}

			outs, err := gateway.GetUnspentOutputs(filters...)
			if err != nil {
				logger.Error("get unspent outputs failed: %v", err)
				wh.Error500(w)
				return
			}

			wh.SendOr404(w, outs)
		}
	}
}

func newWalletSeed(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			logger.Error("new entropy failed when new wallet seed: %v", err)
			wh.Error500(w)
			return
		}

		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			logger.Error("new mnemonic failed when new wallet seed: %v", err)
			wh.Error500(w)
			return
		}

		var rlt = struct {
			Seed string `json:"seed"`
		}{
			mnemonic,
		}

		wh.SendOr404(w, rlt)
	}
}

// RegisterWalletHandlers registers wallet handlers
func RegisterWalletHandlers(mux *http.ServeMux, gateway *daemon.Gateway) {
	// Returns wallet info
	// GET Arguments:
	//      id - Wallet ID.

	//  Gets a wallet .  Will be assigned name if present.
	mux.HandleFunc("/wallet", walletGet(gateway))

	// POST/GET Arguments:
	//		seed [optional]
	//create new wallet
	mux.HandleFunc("/wallet/create", walletCreate(gateway))

	mux.HandleFunc("/wallet/newAddress", walletNewAddresses(gateway))

	// Returns the confirmed and predicted balance for a specific wallet.
	// The predicted balance is the confirmed balance minus any pending
	// spent amount.
	// GET arguments:
	//      id: Wallet ID
	mux.HandleFunc("/wallet/balance", walletBalanceHandler(gateway))

	// Sends coins&hours to another address.
	// POST arguments:
	//  id: Wallet ID
	//  coins: Number of coins to spend
	//  hours: Number of hours to spends
	//  fee: Number of hours to use as fee, on top of the default fee.
	//  Returns total amount spent if successful, otherwise error describing
	//  failure status.
	mux.HandleFunc("/wallet/spend", walletSpendHandler(gateway))

	// GET Arguments:
	//		id: Wallet ID
	// Returns all pending transanction for all addresses by selected Wallet
	mux.HandleFunc("/wallet/transactions", walletTransactionsHandler(gateway))

	// Update wallet label
	// 		GET Arguments:
	// 			id: wallet id
	// 			label: wallet label
	mux.HandleFunc("/wallet/update", walletUpdateHandler(gateway))

	// Returns all loaded wallets
	mux.HandleFunc("/wallets", walletsHandler(gateway))
	// Saves all wallets to disk. Returns nothing if it works. Otherwise returns
	// 500 status with error message.

	// Rescans the wallet directory and loads/unloads wallets based on which
	// files are present. Returns nothing if it works. Otherwise returns
	// 500 status with error message.
	mux.HandleFunc("/wallets/reload", walletsReloadHandler(gateway))

	mux.HandleFunc("/wallets/folderName", getWalletFolder(gateway))

	//get set of unspent outputs
	mux.HandleFunc("/outputs", getOutputsHandler(gateway))

	// get balance of addresses
	mux.HandleFunc("/balance", getBalanceHandler(gateway))

	// generate wallet seed
	mux.Handle("/wallet/newSeed", newWalletSeed(gateway))

	// generate wallet seed
	mux.Handle("/notes", notesHandler(gateway))

	mux.Handle("/notes/create", notesCreate(gateway))
}
