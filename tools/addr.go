package tools

import (
	"github.com/CryptoKass/splashgo/splashkeys"
	"github.com/ethereum/go-ethereum/common"
)

func NewRandomAddress() common.Address {
	priv, _ := splashkeys.GenerateSplashKeys()
	addrBytes := priv.GetAddress()
	return common.BytesToAddress(addrBytes)
}
