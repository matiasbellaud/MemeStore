package handler

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type constants struct {
	memory uint32
	iterations uint32
	keyLength uint32
	hashLength uint32
	parallelism uint8
}

var (
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	ErrInvalidHash = errors.New("The hash is'nt in wrigth format")
	conste = constants {
	memory : 64 * 1024,
	iterations : 4,
	keyLength : 32,
	hashLength : 64,
	parallelism : 3,
	}
)

func HasingFunc(password string) (hash string, err error) {
	key, err := RandomKey(conste.keyLength)
	if err != nil {
		return "",err
	}
	
	hashedPassword := argon2.IDKey([]byte(password), key, conste.iterations, conste.memory, conste.parallelism, conste.hashLength)

	keyb64 := base64.RawStdEncoding.EncodeToString(key)
	hashb64 := base64.RawStdEncoding.EncodeToString(hashedPassword)

	stringToStock := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, conste.memory, conste.iterations, conste.parallelism, keyb64, hashb64)

	return stringToStock, nil
}

func RandomKey(length uint32) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func PasswordEqualHash(password, Hash string) (info bool, err error) {
	key, hash, err := ReadHash(Hash)
	if err != nil {
		return false, err
	}

	secondHash := argon2.IDKey([]byte(password), key, conste.iterations, conste.memory, conste.parallelism, conste.hashLength)

	if subtle.ConstantTimeCompare(hash, secondHash) == 1 {
		return true, nil
	}
	return false, nil
}

func ReadHash(StockHash string) (key, hash []byte, err error) {
	values := strings.Split(StockHash, "$")
	if len(values) != 6 {
		return nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrIncompatibleVersion
	}
	
	key, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, err
	}
	if conste.keyLength != uint32(len(key)) {
		return nil, nil, ErrInvalidHash
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, err
	}
	if conste.hashLength != uint32(len(hash)) {
		return nil, nil, ErrInvalidHash
	}

	return key, hash, nil
}