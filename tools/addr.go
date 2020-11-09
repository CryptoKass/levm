package tools

import (
	"crypto/elliptic"

	"github.com/cryptokass/splashecdsa"
	"github.com/ethereum/go-ethereum/common"
)

func NewRandomAddress() common.Address {
	priv, _ := splashecdsa.GenerateKeys(elliptic.P256())
	addrBytes := priv.GetAddress(true)
	return common.BytesToAddress(addrBytes)
}
