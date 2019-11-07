package ens

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	string2eth "github.com/wealdtech/go-string2eth"
)

type Data struct {
	DomainName     string   `json:"domain_name"`
	SubDomain      string   `json:"subdomain"`
	ToAddress      string   `json:"to_address"`
	Registrant     string   `json:"registrant"`
	Address        string   `json:"address"`
	Contenthash    []byte   `json:"contenthash"`
	TextName       string   `json:"text_name"`
	TextValue      string   `json:"text_value"`
	Abi            string   `json:"abi"`
	AbiName        string   `json:"abi_name"`
	AbiContentType *big.Int `json:"abi_content_type"`
}

var clientUrl = "https://ropsten.infura.io/v3/4a8701b12b674b13b0bd8a332432ac66"

type ResultData struct {
	Name       string `json:"name,omitempty"`
	Domain     string `json:"domain,omitempty"`
	SubDomain  string `json:"subdomain,omitempty"`
	Label      string `json:"label,omitempty"`
	Registrant string `json:"registrant,omitempty"`
	ToAddress  string `json:"to_address,omitempty"`

	// RECORDS
	Address        string   `json:"address,omitempty"`
	Contenthash    []byte   `json:"contenthash,omitempty"`
	TextName       string   `json:"text_name,omitempty"`
	TextValue      string   `json:"text_value,omitempty"`
	Abi            string   `json:"abi,omitempty"`
	AbiName        string   `json:"abi_name,omitempty"`
	AbiContentType *big.Int `json:"abi_content_type,omitempty"`

	Other interface{} `json:"other,omitempty"`
}

type Result struct {
	Status   string      `json:"status"`
	ErrorMsg string      `json:"error_msg,omitempty"`
	Data     *ResultData `json:"data,omitempty"`
}

func int2big(i string) *big.Int {
	bi := new(big.Int)
	_, err := fmt.Sscan(i, bi)
	if err != nil {
		i2 := "0"
		fmt.Sscan(i2, bi)
		return bi
	} else {
		return bi
	}
}

func generateTxOpts(client *ethclient.Client, sender common.Address, valueStr string) (*bind.TransactOpts, error) {
	// key, err := crypto.HexToECDSA(os.Getenv(fmt.Sprintf("PRIVATE_KEY_%x", sender)))
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to obtain private key for %x", sender)
	// }
	key, err := crypto.HexToECDSA("D987F6B649B04D427E7AA188662BF183A04388114E84B5E00986F3E265D0905B")

	signer := keySigner(big.NewInt(3), key)
	if signer == nil {
		return nil, errors.New("no signer")
	}

	value, err := string2eth.StringToWei(valueStr)
	if err != nil {
		return nil, err
	}

	curNonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return nil, err
	}
	nonce := int64(curNonce)

	opts := &bind.TransactOpts{
		From:     sender,
		Signer:   signer,
		GasPrice: big.NewInt(10000000000),
		Value:    value,
		Nonce:    big.NewInt(0).SetInt64(nonce),
	}

	return opts, nil
}

func keySigner(chainID *big.Int, key *ecdsa.PrivateKey) bind.SignerFn {
	return func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
}

func waitForTransaction(client *ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err := client.TransactionByHash(context.Background(), txHash)
		if err == nil && !pending {
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func hasPrivateKey(address common.Address) bool {
	_, err := crypto.HexToECDSA(os.Getenv(fmt.Sprintf("PRIVATE_KEY_%x", address)))
	return err == nil
}
