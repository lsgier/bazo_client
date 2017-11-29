package client

import (
	"fmt"
	"github.com/mchetelat/bazo_miner/p2p"
	"github.com/mchetelat/bazo_miner/protocol"
	"log"
	"os"
)

var (
	err     error
	msgType uint8
	tx      protocol.Transaction
	logger  *log.Logger
)

const (
	USAGE_MSG = "Usage: bazo_client [pubKey|accTx|fundsTx|configTx] ...\n"
)

func Init() {
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	initState()
}

func State(keyFile string) {
	pubKeyTmp, _, err := extractKeyFromFile(keyFile)

	var pubKey [64]byte
	copy(pubKey[:32], pubKeyTmp.X.Bytes())
	copy(pubKey[32:], pubKeyTmp.Y.Bytes())

	if err != nil {
		fmt.Printf("%v\n%v", err, USAGE_MSG)
	} else {
		fmt.Printf("My address: %x\n", pubKey)

		acc, err := GetAccount(pubKey)
		if err != nil {
			logger.Println(err)
		} else {
			logger.Printf(acc.String())
		}
	}
}

func Process(args []string) {
	switch args[0] {
	case "accTx":
		tx, err = parseAccTx(os.Args[2:])
		msgType = p2p.ACCTX_BRDCST
	case "fundsTx":
		tx, err = parseFundsTx(os.Args[2:])
		msgType = p2p.FUNDSTX_BRDCST
	case "configTx":
		tx, err = parseConfigTx(os.Args[2:])
		msgType = p2p.CONFIGTX_BRDCST
	default:
		fmt.Printf("Usage: bazo_client [accTx|fundsTx|configTx] ...\n")
		return
	}

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if err := SendTx(tx, msgType); err != nil {
		logger.Printf("%v\n", err)
	} else {
		logger.Printf("Successfully sent the following tansaction:%v", tx)
	}
}
