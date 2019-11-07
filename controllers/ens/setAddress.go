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

func setAddress(data Data) Result {
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

	resSet, err := resolver.SetAddress(opts, common.HexToAddress(data.ToAddress))
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	resultData.Name = data.DomainName
	resultData.Registrant = data.Registrant
	resultData.ToAddress = data.ToAddress
	resultData.Other = &resSet
	result.Data = &resultData
	return result
}

func SetAddress(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(setAddress(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
