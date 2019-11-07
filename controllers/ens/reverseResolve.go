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

func reverseResolve(data Data) Result {
	var result Result
	var resultData ResultData

	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	reverse, err := ens.ReverseResolve(client, common.HexToAddress(data.Address))
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	if reverse == "" {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprintf("%s has no reverse lookup", data.Address)
		return result
	}

	result.Status = "success"
	resultData.Address = data.Address
	resultData.Name = reverse
	result.Data = &resultData
	return result
}

func ReverseResolve(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(reverseResolve(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
