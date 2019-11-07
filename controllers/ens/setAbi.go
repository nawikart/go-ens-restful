package ens

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens"
)

func setAbi(data Data) Result {
	var result Result
	var resultData ResultData
	registrant := common.HexToAddress(data.Registrant)

	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	opts, err := generateTxOpts(client, registrant, "0")
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	resolver, err := ens.NewResolver(client, data.DomainName)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	resSet, err := resolver.SetABI(opts, data.AbiName, data.Abi, data.AbiContentType)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	resultData.Name = data.DomainName
	resultData.Registrant = data.Registrant
	resultData.AbiName = data.AbiName
	resultData.Abi = data.Abi
	resultData.AbiContentType = data.AbiContentType
	resultData.Other = &resSet
	result.Data = &resultData
	return result
}

func SetABI(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(setAbi(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
