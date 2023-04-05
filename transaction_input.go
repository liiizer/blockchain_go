package main

import "bytes"

// 交易输入数据结构
type TXInput struct {
	Txid      []byte `json:"txid"`
	Vout      int    `json:"vout"`
	Signature []byte `json:"signature"`
	PubKey    []byte `json:"pubkey"`
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
