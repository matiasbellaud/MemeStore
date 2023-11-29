package Hash

import (
	"crypto/sha256"
	"fmt"
)

func HashPassord(password string) string{
	AreyHash := sha256.Sum256([]byte("hello world\n"))
	Hash := fmt.Sprintf("%x", AreyHash)
	return Hash
}

func IsGoodPassword(entry string, hashPassword string) bool {
	return HashPassord(entry) == hashPassword
}