package npub

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/bech32"
)

// Decode bech32 public key with 'npub' human-readable part into hex encoded public key
func Decode(npub string) (string, error) {
	hrp, byt, err := bech32.DecodeNoLimit(npub)
	if err != nil {
		return hrp, err
	}
	grp, err := bech32.ConvertBits(byt, 5, 8, false)
	if err != nil {
		return hrp, err
	}
	if len(grp) < 32 {
		return hrp, fmt.Errorf("invalid npub")
	}
	return hex.EncodeToString(grp[0:32]), nil
}

// Encode hex public key as bech32 with "npub" human readable part
func Encode(pubKeyHex string) (string, error) {
	str, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid hex encoded public key")
	}
	grp, err := bech32.ConvertBits(str, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode("npub", grp)
}

func New() (prvKeyHex string, pubKeyHex string, npub string, err error) {
	prvKey, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", "", err
	}
	prvKeyHex = hex.EncodeToString(prvKey.Serialize())
	pubKeyHex = hex.EncodeToString(prvKey.PubKey().SerializeCompressed()[1:])
	npub, err = Encode(pubKeyHex)
	if err != nil {
		return "", "", "", err
	}
	return prvKeyHex, pubKeyHex, npub, nil
}
