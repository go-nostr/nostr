package subscriptionid

import (
	"fmt"
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func New() string {
	b := make([]rune, 64)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func Validate(subscriptionID string) error {
	if subscriptionID == "" || len(subscriptionID) > 64 {
		return fmt.Errorf("invalid subscription id")
	}
	return nil
}
