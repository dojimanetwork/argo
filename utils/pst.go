package utils

import "github.com/dojimanetwork/argo/types"


const (
	ArDrive = "-8A6RexFkpfWwuyVO98wzSFZh0d6VJuI-buTJvlwOJQ"
	Verto = "usjm4PCxUd5mtaon7zc97-dt-3qf67yPyqgzLnLqk5A"
	ArVerify = "f6lW-sKxsc340p8eBBL2i_fnmSI_fRSFmkqvzqyUsRs"
)

var (
 supportedPstContracts = [3]string{ArDrive, Verto, ArVerify}
)


func PstTransferTags(contractId string, target string, qty int64) ([]types.Tag, error) {
	input := types.Input{
		"function": "transfer",
		"target":   target,
		"qty":      qty,
	}

	inputStr, err := input.ToString()
	if err != nil {
		return nil, err
	}

	pstTags := []types.Tag{
		{Name: "App-Name", Value: "SmartWeaveAction"},
		{Name: "App-Version", Value: "0.3.0"},
		{Name: "Contract", Value: contractId},
		{Name: "Input", Value: inputStr},
	}
	return pstTags, nil
}
