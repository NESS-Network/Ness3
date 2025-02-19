/*
ness-daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
    "github.com/MDLlife/MDL/src/fiber"
	"os"

	"github.com/MDLlife/MDL/src/readable"
	"github.com/MDLlife/MDL/src/mdl"
	"github.com/MDLlife/MDL/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
<<<<<<< HEAD:cmd/ness3/ness3.go
	Version = "0.27.1"
=======
    Version = "0.27.1"
>>>>>>> release/0.27.1:cmd/mdl/mdl.go
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// CoinName name of coin
	CoinName = "ness3"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "05d4045854103f8a8938bb701cc4101c38942a180ba02d328d6f880bf37b387c47e95813d061f94bdf5d894bfebf17f933c5fc92fc9d010480765257c3d19d9b00"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "24GJTLPMoz61sV4J4qg1n14x5qqDwXqyJJy"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02933015bd2fa1e0a885c05fb08eb7c647bf8c3188ed5120b51d0d09ccaf525036"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1426562704
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 200000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
		"192.243.100.192:6660",
		"167.114.97.165:6660",
		"198.100.144.39:6660",
		"94.23.56.111:6660",
	}

<<<<<<< HEAD:cmd/ness3/ness3.go
	nodeConfig = mdl.NewNodeConfig(ConfigMode, mdl.NodeParameters{
		CoinName:                       CoinName,
		GenesisSignatureStr:            GenesisSignatureStr,
		GenesisAddressStr:              GenesisAddressStr,
		GenesisCoinVolume:              GenesisCoinVolume,
		GenesisTimestamp:               GenesisTimestamp,
		BlockchainPubkeyStr:            BlockchainPubkeyStr,
		BlockchainSeckeyStr:            BlockchainSeckeyStr,
		DefaultConnections:             DefaultConnections,
		PeerListURL:                    "http://nodes.privateness.network/blockchain/peers.txt",
		Port:                           6660,
		WebInterfacePort:               6420,
		DataDirectory:                  "$HOME/.ness3",
=======
	nodeConfig = mdl.NewNodeConfig(ConfigMode, fiber.NodeConfig{
		CoinName:            CoinName,
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "",
		Port:                7800,
		WebInterfacePort:    8320,
		DataDirectory:       "$HOME/.mdl",

        UnconfirmedBurnFactor:          10,
        UnconfirmedMaxTransactionSize:  32768,
        UnconfirmedMaxDropletPrecision: 3,
        CreateBlockBurnFactor:          10,
        CreateBlockMaxTransactionSize:  32768,
        CreateBlockMaxDropletPrecision: 3,
        MaxBlockTransactionsSize:       32768,

        DisplayName:           "MDL Talent Hub",
        Ticker:                "MDL",
        CoinHoursName:         "Talent Hours",
        CoinHoursNameSingular: "Talent Hour",
        CoinHoursTicker:       "MTH",
        ExplorerURL:           "https://explorer.mdl.wtf",
        VersionURL:            "https://api.github.com/repos/MDLlife/MDL/tags",
        Bip44Coin:             8000,
>>>>>>> release/0.27.1:cmd/mdl/mdl.go
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := mdl.NewCoin(mdl.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}
