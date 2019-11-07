package ens

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens"
)

func resolve(data Data) Result {
	var result Result
	var resultData ResultData

	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	address, err := ens.Resolve(client, data.DomainName)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"
	resultData.Name = data.DomainName
	resultData.Address = fmt.Sprint(address.Hex())
	result.Data = &resultData
	return result
}

func Resolve(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(resolve(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
