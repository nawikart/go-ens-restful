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

func createSubdomain(data Data) Result {
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

	name, err := ens.NewName(client, data.DomainName)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	tx, err := name.CreateSubdomain(data.SubDomain, registrant, opts)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}
	waitForTransaction(client, tx.Hash())

	// // Confirm registrantship of the subdomain
	subdomain := fmt.Sprintf("%s.%s", data.SubDomain, data.DomainName)

	registry, err := ens.NewRegistry(client)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	// controller, err := registry.Owner(subdomain)
	_, err = registry.Owner(subdomain)
	if err != nil {
		result.Status = "failed"
		result.ErrorMsg = fmt.Sprint(err)
		return result
	}

	result.Status = "success"

	resultData.Name = data.DomainName
	resultData.SubDomain = data.SubDomain
	resultData.Registrant = data.Registrant
	result.Data = &resultData
	return result
}

func CreateSubdomain(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var data Data
	if json.Unmarshal(body, &data) == nil {
		json.NewEncoder(w).Encode(createSubdomain(data))
		return
	}
	http.Error(w, "invalid_json_data", http.StatusNotAcceptable)
}
