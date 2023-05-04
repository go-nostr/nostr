package npub

import (
	"encoding/hex"
	"fmt"

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
		return hrp, fmt.Errorf("byt is less than 32 bytes (%d)", len(grp))
	}
	return hex.EncodeToString(grp[0:32]), nil
}

// Encode hex encoded public key as bech32 with "npub" human readable part
func Encode(pubKeyHex string) (string, error) {
	str, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key hex: %w", err)
	}
	grp, err := bech32.ConvertBits(str, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode("npub", grp)
}
