package argo

import (
	"errors"
	"strconv"

	"github.com/dojimanetwork/argo/types"
	"github.com/dojimanetwork/argo/utils"
)

func (w *Wallet) CreateAndSignBundleItem(data []byte, signatureType int, target string, anchor string, tags []types.Tag) (di types.BundleItem, err error) {
	bundleItem := utils.NewBundleItem(w.Owner(), strconv.Itoa(signatureType), target, anchor, data, tags)
	// sign
	err = utils.SignBundleItem(bundleItem, w.Signer.PrvKey)
	if err != nil {
		return di, err
	}
	err = utils.GenerateItemBinary(bundleItem)
	if err != nil {
		return di, err
	}
	return *bundleItem, nil
}

func (w *Wallet) SendBundleTxSpeedUp(bundleBinary []byte, tags []types.Tag, txSpeed int64) (types.Transaction, error) {
	bundleTags := []types.Tag{
		{Name: "Bundle-Format", Value: "binary"},
		{Name: "Bundle-Version", Value: "2.0.0"},
	}
	// check tags cannot include bundleTags Name
	mmap := map[string]struct{}{
		"Bundle-Format":  {},
		"Bundle-Version": {},
	}
	for _, tag := range tags {
		if _, ok := mmap[tag.Name]; ok {
			return types.Transaction{}, errors.New("tags can not set bundleTags")
		}
	}

	txTags := make([]types.Tag, 0)
	txTags = append(bundleTags, tags...)
	return w.SendDataSpeedUp(bundleBinary, txTags, txSpeed)
}

func (w *Wallet) SendBundleTx(bundleBinary []byte, tags []types.Tag) (types.Transaction, error) {
	return w.SendBundleTxSpeedUp(bundleBinary, tags, 0)
}
