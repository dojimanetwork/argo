package argo

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testWallet *Wallet
var err error

func init() {
	clientUrl := "https://arweave.net"
	testWallet, err = NewWallet([]byte(`{
		"d": "",
		"dp": "",
		"dq": "",
		"e": "AQAB",
		"ext": true,
		"kty": "RSA",
		"n":"xfEQqmKNrMoHg6v3U599Yzs61k8hQwcTZyOXMjItW8W30ay2PWdbh7jnkNVkKtd_AmxoeGt-ixEGN2FDhLYmkc06ehez5faN-sppup7-OZPJ3Dz6BODswtH7xd_cPvGQ0tzsUtZ7jDHFKpGtifi4O3VcbZ8VW113DRYVHDL-wNuT_xWVFJVV-MiJJ45yFWD9_aQbrTz1DFQNT_LZOlR_wOhkTYmZ1OoZjF6tPZK9owm2pxqo18tHjjPjwXZ4rx7sFYpL7UdrYQ9JCvuDoGEZ2dhvrxjspB1tOsUllwjIJTFEBt5A7our6ra9j5J1XdTwxRPj2m3pQmatPnyURZxT4NoxPpfAB4UE30clz0LnnSRHYPAEkdrTI95ovBum5trffzxwOcq6bMnP8uunQvphz--TCYhfpEOwb2_4_bP23UaL45a9ZNQZM9x-6q5lWshpMHMNdwcgNclm2hPIJ6xVYTek1w-O5cGBWlg_hYiU9Dw6-TZOLVgImt6hK69v9lJr7vkxWi_iglknKs3wZ9JhC0uHIm3hJfAfoCSxvyMypEfNFdqSS9HgiyEq--BRSI_MGyBro0p9INqIf1Re-9jgXz4E2LO3n8EVnaJkeCscaT1EO2mY7RNqIKUTwEf8WGv_2cChmUQdf2vTHB1nJ2BQuxZomQniGc3shscnTpwg-_c",
		"p": "",
		"q": "",
		"qi": ""
	  }`),
		clientUrl)

	if err != nil {
		panic(err)
	}
}

func TestPubKey(t *testing.T) {
	pubKey := testWallet.PubKey
	assert.Equal(t, "nQ9iy1fRM2xrgggjHhN1xZUnOkm9B4KFsJzH70v7uLMVyDqfyIJEVXeJ4Jhk_8KpjzYQ1kYfnCMjeXnhTUfY3PbeqY4PsK5nTje0uoOe1XGogeGAyKr6mVtKPhBku-aq1gz7LLRHndO2tvLRbLwX1931vNk94bSfJPYgMfU7OXxFXbTdKU38W6u9ShoaJGgUQI1GObd_sid1UVniCmu7P-99XPkixqyacsrkHzBajGz1S7jGmpQR669KWE9Z0unvH0KSHxAKoDD7Q7QZO7_4ujTBaIFwy_SJUxzVV8G33xvs7edmRdiqMdVK5W0LED9gbS4dv_aee9IxUJQqulSqZphPgShIiGNl9TcL5iUi9gc9cXR7ISyavos6VGiem_A-S-5f-_OKxoeZzvgAQda8sD6jtBTTuM5eLvgAbosbaSi7zFYCN7zeFdB72OfvCh72ZWSpBMH3dkdxsKCDmXUXvPdDLEnnRS87-MP5RV9Z6foq_YSEN5MFTMDdo4CpFGYl6mWTP6wUP8oM3Mpz3-_HotwSZEjASvWtiff2tc1fDHulVMYIutd52Fis_FKj6K1fzpiDYVA1W3cV4P28Q1-uF3CZ8nJEa5FXchB9lFrXB4HvsJVG6LPSt-y2R9parGi1_kEc6vOYIesKspgZ0hLyIKtqpTQFiPgKRlyUc-WEn5E", base64.RawURLEncoding.EncodeToString(pubKey.N.Bytes()))
}

func TestAddress(t *testing.T) {
	addr := testWallet.Address
	assert.Equal(t, "Y3agEpiSLqDbRVID7aGUkuLw8G6qGEgJ6MBYUVc-ADA", addr)
}

// test sand ar without data
func TestWallet_SendAR(t *testing.T) {
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("../example/testKey.json", arNode) // your wallet private key
	// assert.NoError(t, err)
	//
	// target := "Goueytjwney8mRqbWBwuxbk485svPUWxFQojteZpTx8"
	// amount := big.NewFloat(0.001)
	// tags := []types.Tag{
	// 	{Name: "argo", Value: "sendAR"},
	// }
	// id,  err := w.SendAR(amount, target, tags)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s \n", id)
}

// test send small size file
func TestWallet_SendDataSpeedUp01(t *testing.T) {
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("./example/testKey.json", arNode) // your wallet private key
	// assert.NoError(t, err)
	//
	// // data := []byte("aaa this is a argo test small size file data") // small file
	// data := make([]byte, 255*1024)
	// for i := 0; i < len(data); i++ {
	// 	data[i] = byte('b' + i)
	// }
	// tags := []types.Tag{
	// 	{Name: "argo", Value: "SMDT"},
	// }
	// id, err := w.SendDataSpeedUp(data, tags, 50)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s", id)
}

// test send big size file
func TestWallet_SendDataSpeedUp02(t *testing.T) {
	// proxyUrl := "http://127.0.0.1:8001"
	// arNode := "https://arweave.net"
	// w, err := NewWalletFromPath("./wallet/account1.json", arNode, proxyUrl) // your wallet private key
	// assert.NoError(t, err)
	//
	// data, err := ioutil.ReadFile("/Users/sandyzhou/Downloads/abc.jpeg")
	// if err != nil {
	// 	panic(err)
	// }
	// tags := []types.Tag{
	// 	{Name: "Sender", Value: "Jie"},
	// 	{Name: "Data-Introduce", Value: "Happy anniversary, my google and dearest! I‘m so grateful to have you in my life. I love you to infinity and beyond! (⁎⁍̴̛ᴗ⁍̴̛⁎)"},
	// }
	// id, err := w.SendDataSpeedUp(data, tags, 10)
	// assert.NoError(t, err)
	// t.Logf("tx hash: %s", id)
}

func Test_SendPstTransfer(t *testing.T) {
	// w, err := NewWalletFromPath("./wallet/account1.json","https://arweave.net")
	// assert.NoError(t, err)
	//
	// contractId := "usjm4PCxUd5mtaon7zc97-dt-3qf67yPyqgzLnLqk5A"
	// target := "Ii5wAMlLNz13n26nYY45mcZErwZLjICmYd46GZvn4ck"
	// qty := int64(1)
	// arId, err := w.SendPstTransfer(contractId,target,qty,nil,50)
	// assert.NoError(t, err)
	// t.Log(arId)
}
