package argo

import (
	"fmt"
	"testing"

	"github.com/dojimanetwork/argo/utils"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"fmt"
// 	"testing"

// func TestGetTransactionByID(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	fmt.Println(client.GetTransactionByID("FgcKlptyDXSgEonYfy5cNBimq7GJ4h8h6L6pxuuYOBc"))
// }

// func TestGetTransactionPrice(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	target := ""
// 	reward, err := client.GetTransactionPrice([]byte("123"), &target)
// 	assert.NoError(t, err)
// 	fmt.Println(reward)
// }

// func TestGetLastTransactionID(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	lastTx, err := client.GetLastTransactionID("dQzTM9hXV5MD1fRniOKI3MvPF_-8b2XDLmpfcMN9hi8")
// 	assert.NoError(t, err)
// 	fmt.Println(lastTx)
// }

// func TestGetTransactionAnchor(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	fmt.Println(client.GetTransactionAnchor())
// }

// func TestSubmitTransaction(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	fmt.Println(
// 		client.SubmitTransaction(&types.Transaction{
// 			ID: "n1iKT3trKn6Uvd1d8XyOqKBy8r-8SSBtGA62m3puK5k",
// 		}),
// 	)
// }

// func TestArql(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	fmt.Println(
// 		client.Arql(`
// 		{
// 			"op": "and",
// 			"expr1": {
// 				"op": "equals",
// 				"expr1": "TokenSymbol",
// 				"expr2": "ROL"
// 			},
// 			"expr2": {
// 				"op": "equals",
// 				"expr1": "CreatedBy",
// 				"expr2": "dQzTM9hXV5MD1fRniOKI3MvPF_-8b2XDLmpfcMN9hi8"
// 			}
// 		}
// 		`),
// 	)
// }

// func TestGraphQL(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	data, err := client.GraphQL(`
// 	{
// 		transactions(
// 			tags: [
// 					{
// 							name: "TokenSymbol",
// 							values: "ROL"
// 					},
// 			]
// 			sort: HEIGHT_ASC
// 		) {
// 			edges {
// 				node {
// 					id
// 					tags {
// 						name
// 						value
// 					}
// 				}
// 			}
// 		}
// 	}`)
// 	assert.NoError(t, err)
// 	t.Log(string(data))
// }

// func TestGetWalletBalance(t *testing.T) {
// 	client := NewClient("https://arweave.net")
// 	fmt.Println(
// 		client.GetWalletBalance("dQzTM9hXV5MD1fRniOKI3MvPF_-8b2XDLmpfcMN9hi8"),
// 	)
// }

func TestClient_DownloadChunkData(t *testing.T) {
	// client := NewClient("https://arweave.net")
	// id := "ybEmme6TE3JKwnSYciPCjnAINwi_CWthomsxBes-kYk"
	// data, err := client.GetTransactionData(id, "jpg")
	// assert.NoError(t, err)
	//
	// t.Log(len(data))
	// err = ioutil.WriteFile("photo.jpg", data, 0777)
	// assert.NoError(t, err)
}

func TestClient_Arql(t *testing.T) {
	// client := NewClient("https://arweave.dev")
	// id := "PvLGaQzn9MOwucO91uuMGRnq8pj1qlwbURPqhmW0UiM"
	//
	// status, err := client.GetTransactionStatus(id)
	// assert.NoError(t, err)
	// t.Log(status)
}

func TestClient_VerifyTx(t *testing.T) {
	// txId := "XOzxw5kaYJrt9Vljj23pA5_6b63kY2ydQ0lPfnhksMA"
	txId := "_fVj-WyEtXV3URXlNkSnHVGupl7_DM1UWZ64WMdhPkU"
	client := NewClient("https://arweave.net")
	tx, err := client.GetTransactionByID(txId)
	assert.NoError(t, err)
	t.Log(tx.Format)
	t.Log(utils.TagsDecode(tx.Tags))
	err = utils.VerifyTransaction(*tx)
	assert.NoError(t, err)
}

func TestGetTransaction(t *testing.T) {
	arNode := "https://arweave.net"
	cli := NewClient(arNode)

	// on chain tx
	txId := "ggt-x5Q_niHifdNzMxZrhiibKf0KQ-cJun0UIBBa-yA"
	txStatus, err := cli.GetTransactionStatus(txId)
	assert.NoError(t, err)
	assert.Equal(t, 575660, txStatus.BlockHeight)
	tx, err := cli.GetTransactionByID(txId)
	assert.NoError(t, err)
	assert.Equal(t, "0pu7-Otb-AH6SSSX_rfUmpTkwh3Nmhpztd_IT8nYXDwBE6P3B-eJSBuaTBeLypx4", tx.LastTx)

	// not exist tx
	txId = "KPlEyCrcs2rDHBFn2f0UUn2NZQKfawGb_EnBfip8ayA"
	txStatus, err = cli.GetTransactionStatus(txId)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, txStatus)
	tx, err = cli.GetTransactionByID(txId)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, tx)

	// // pending tx
	// txId = "muANv_lsyZKC5C8fTxQaC2dCCyGDao8z35ECuGdIBP8" // need send a new tx create pending status
	// txStatus, err = cli.GetTransactionStatus(txId)
	// assert.Equal(t, "Pending",err.Error())
	// assert.Nil(t, txStatus)
	// tx, err = cli.GetTransactionByID(txId)
	// assert.Equal(t, "Pending",err.Error())
	// assert.Nil(t, txStatus)
}

func TestClient_GetTransactionTags(t *testing.T) {
	arNode := "https://arweave.net"
	cli := NewClient(arNode)
	id := "gdXUJuj9EZm99TmeES7zRHCJtnJoP3XgYo_7KJNV8Vw"
	tags, err := cli.GetTransactionTags(id)
	assert.NoError(t, err)
	assert.Equal(t, "App", tags[0].Name)
	assert.Equal(t, "Version", tags[1].Name)
	assert.Equal(t, "Owner", tags[2].Name)
}

func TestClient_GetBlockByHeight(t *testing.T) {
	arNode := "https://arweave.net"
	cli := NewClient(arNode)
	block, err := cli.GetBlockByHeight(793791)
	fmt.Printf("block %v", block)
	assert.NoError(t, err)
	assert.Equal(t, "ci2uJhYmdldgkHbScDClCwAA0eqn7dCduAEpLfRorSA", block.Nonce)
}

func TestClient_GetTransactionDataByGateway(t *testing.T) {
	arNode := "https://ar-test.h4s.dojima.network"
	cli := NewClient(arNode)
	err := cli.MintTestAr("2txTDSdb_RjG12uHZlVsB5jrfPzqxtzScKTtPef2KZ0", "100000000000000000000")
	bal, err := cli.GetWalletBalance("2txTDSdb_RjG12uHZlVsB5jrfPzqxtzScKTtPef2KZ0")
	fmt.Println(bal)
	assert.NoError(t, err)
}

func TestClient_GetPeers(t *testing.T) {
	arNode := "https://arweave.net"
	cli := NewClient(arNode)
	peers, err := cli.GetPeers()
	assert.NoError(t, err)
	t.Log(len(peers))
}

func Test_GetTxDataFromPeers(t *testing.T) {
	cli := NewClient("https://arweave.net")
	txId := "J5FY1Ovd6JJ49WFHfCf-1wDM1TbaPSdKnGIB_8ePErE"
	data, err := cli.GetTxDataFromPeers(txId)

	assert.NoError(t, err)

	assert.NoError(t, err)
	t.Log(len(data))

	// verify data root
	chunks := utils.GenerateChunks(data)
	dataRoot := utils.Base64Encode(chunks.DataRoot)
	tx, err := cli.GetTransactionByID(txId)
	assert.NoError(t, err)
	assert.Equal(t, tx.DataRoot, dataRoot)
}

func TestClient_BroadcastData(t *testing.T) {
	cli := NewClient("https://arweave.net")
	txId := "J5FY1Ovd6JJ49WFHfCf-1wDM1TbaPSdKnGIB_8ePErE"
	data, err := cli.GetTransactionData(txId, "json")
	assert.NoError(t, err)

	err = cli.BroadcastData(txId, data, 20)
	assert.NoError(t, err)
}

func TestClient_GetBlockFromPeers(t *testing.T) {
	cli := NewClient("https://arweave.net")
	block, err := cli.GetBlockFromPeers(793755)
	assert.NoError(t, err)
	t.Log(block.Txs)
}

func TestClient_GetTxFromPeers(t *testing.T) {
	cli := NewClient("https://arweave.net")
	arId := "5MiJDf2gFh4w3RXs1iXRrM9V8UwtnxX6xFATgxUqUN4"
	tx, err := cli.GetTxFromPeers(arId)
	assert.NoError(t, err)
	t.Log(tx)
}

func TestClient_GetUnconfirmedTx(t *testing.T) {
	cli := NewClient("https://arweave.net")
	arId := "5MiJDf2gFh4w3RXs1iXRrM9V8UwtnxX6xFATgxUqUN4"
	tx, err := cli.GetUnconfirmedTx(arId)
	assert.NoError(t, err)
	t.Log(tx)
}

func TestClient_GetUnconfirmedTxFromPeers(t *testing.T) {
	cli := NewClient("https://arweave.net")
	arId := "5MiJDf2gFh4w3RXs1iXRrM9V8UwtnxX6xFATgxUqUN4"
	tx, err := cli.GetUnconfirmedTxFromPeers(arId)
	assert.NoError(t, err)
	t.Log(tx)
}

func TestNewClient(t *testing.T) {
	cli := NewClient("https://arweave.net")
	res, err := cli.GetPendingTxIds()
	assert.NoError(t, err)
	t.Log("pending tx number:", len(res))
}

func TestParseOwner(t *testing.T) {
	owner, _ := utils.Base64Decode("3iX6z3CrZsp-bqqb-bRBaz9AwGZEnrcB5_0yLgdxL9NIAuSTMmvT3_lp8aC9X8JW17MvbsBEp6PjSRiDG2vg1RQvV4pxhBNM9hiGf_56pEqO5gpEJzAyJWPXLHyHAFBeNoRzYf_63V4_s0YQbS74WOn7iVb_0l4oxmXkjWZEBmjFH4x-L6T_jkW14i8YGy3s5sZ4qU-N_A-tU8dmWqUabrlm-bOnN3HnYKmKhKxz0-lWmD9WP6FG0oK8ZVnbWZgDyBS7XhxgZ5ZDqOXhYo70ztZKixhGMIZ4V2CyLJSbe6Z1jadfQyIUxXvLvPLgnkmKga16HszIGzAonDDJN-keSkxcJsVY1UH9LZbBqZoadj2-q-1wCOJ2qc88C4yZGwNJtV3uc8Z5DIW8sWv6u_m3tuie74PH3Ac8eZ9iprJAkPiYKy5llhEGB0bd60q3aiA5J-tAVHApXZxgfSph6JGSkg5WQ4jlqPutFMexhg-uIM9BKOJtm0p3d1xCZ8pXFEyZWjftJPa5xgBBBWQfuRy6qikLpuul7oISPgCtZHZKI7-80ioUCzQfYrs31l3RXlvKuHiN5DPERhQcqSy7VN1lPAfAt2V4bNY4LOQ_NO4qV8e9QCeJp5vrXKrVW9ig25HmcHhucRRSImtEoFIXTCWCIbV80dbCXXHtFENFNxMhY1E is invalid sender address: address format not supported: 3iX6z3CrZsp-bqqb-bRBaz9AwGZEnrcB5_0yLgdxL9NIAuSTMmvT3_lp8aC9X8JW17MvbsBEp6PjSRiDG2vg1RQvV4pxhBNM9hiGf_56pEqO5gpEJzAyJWPXLHyHAFBeNoRzYf_63V4_s0YQbS74WOn7iVb_0l4oxmXkjWZEBmjFH4x-L6T_jkW14i8YGy3s5sZ4qU-N_A-tU8dmWqUabrlm-bOnN3HnYKmKhKxz0-lWmD9WP6FG0oK8ZVnbWZgDyBS7XhxgZ5ZDqOXhYo70ztZKixhGMIZ4V2CyLJSbe6Z1jadfQyIUxXvLvPLgnkmKga16HszIGzAonDDJN-keSkxcJsVY1UH9LZbBqZoadj2-q-1wCOJ2qc88C4yZGwNJtV3uc8Z5DIW8sWv6u_m3tuie74PH3Ac8eZ9iprJAkPiYKy5llhEGB0bd60q3aiA5J-tAVHApXZxgfSph6JGSkg5WQ4jlqPutFMexhg-uIM9BKOJtm0p3d1xCZ8pXFEyZWjftJPa5xgBBBWQfuRy6qikLpuul7oISPgCtZHZKI7-80ioUCzQfYrs31l3RXlvKuHiN5DPERhQcqSy7VN1lPAfAt2V4bNY4LOQ_NO4qV8e9QCeJp5vrXKrVW9ig25HmcHhucRRSImtEoFIXTCWCIbV80dbCXXHtFENFNxMhY1E")
	t.Log("owner", owner)
}
