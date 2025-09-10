package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateShortID(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:8]
}
