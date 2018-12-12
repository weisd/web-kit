package istring

import (
	"crypto/sha1"
	"encoding/base64"
)

// HashPassword HashPassword
func HashPassword(pass, salt string) string {
	h := sha1.New()
	h.Write([]byte(pass + salt))
	checksum := h.Sum(nil)
	passHashed := base64.StdEncoding.EncodeToString([]byte(string(checksum) + salt))
	return passHashed
}
