package nsec

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/bech32"
)

// Decode bech32 public key with 'nsec' human-readable part into hex encoded private key
func Decode(nsec string) (string, error) {
	hrp, byt, err := bech32.DecodeNoLimit(nsec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if hrp != "nsec" {
		return "", fmt.Errorf("invalid nsec")
	}
	grp, err := bech32.ConvertBits(byt, 5, 8, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hex.EncodeToString(grp[0:32]), nil
}

// Encode hex private key as bech32 with "nsec" human readable part
func Encode(prvKeyHex string) (string, error) {
	data, err := hex.DecodeString(prvKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid hex encoded private key")
	}
	grp, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode("nsec", grp)
}

func New() (prvKeyHex string, pubKeyHex string, nsec string, err error) {
	prvKey, err := btcec.NewPrivateKey()
	if err != nil {
		return "", "", "", err
	}
	prvKeyHex = hex.EncodeToString(prvKey.Serialize())
	pubKeyHex = hex.EncodeToString(prvKey.PubKey().SerializeCompressed()[1:])
	nsec, err = Encode(prvKeyHex)
	if err != nil {
		return "", "", "", err
	}
	return prvKeyHex, pubKeyHex, nsec, nil
}
