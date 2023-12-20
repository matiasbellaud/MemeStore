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

// initialasing the structure used for hashing
type constants struct {
	memory uint32
	iterations uint32
	keyLength uint32
	hashLength uint32
	parallelism uint8
}

// set the var use in all folowing func, first for errors second to the hashing method
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
	/*
	principal function to hash pasword
	-------------------------------------------------------
	input : the password a string
	output : the formated hash (version,key,hash) and potential error
	-------------------------------------------------------
	at first the function générate a key to make the argon2 hash
	if the key generation faild the function return an empty string and the error

	after that func hash the password and then encode in Base64 the key and the hash
	finaly function create a string who contain the argon version, the base 64 encripted key and base64 encrypted hash

	at the end return a final string and a nil error
	*/
	key, err := RandomKey(conste.keyLength)
	if err != nil {
		return "",err
	}
	
	hashedPassword := argon2.IDKey([]byte(password), key, conste.iterations, conste.memory, conste.parallelism, conste.hashLength)

	keyb64 := base64.RawStdEncoding.EncodeToString(key)
	hashb64 := base64.RawStdEncoding.EncodeToString(hashedPassword)

	stringToStock := fmt.Sprintf("$v=%d$%s$%s", argon2.Version, keyb64, hashb64)

	return stringToStock, nil
}

func RandomKey(length uint32) ([]byte, error) {
	/*
	function who generate random key for argon2 hash
	-------------------------------------------------
	input : an unit 32 represent the lenght of the key
	output : the key in byte format, potential error
	-------------------------------------------------
	first step the func create the array of byte whith good length
	next it randomise the content of the array
	if the randomising fail func return a nil array and the error

	at the end it return the key in array of byte form and a nil error
	*/
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func PasswordEqualHash(password, Hash string) (info bool, err error) {
	/*
	function who compare a password to a stocked hash
	---------------------------------------------------
	input : the password and stocked hash the both are string
	output : the boolean born from the comparing and possible error
	---------------------------------------------------
	first the func use the ReadHash func
	if the ReadHash return an error the func return false and the error
	
	after the func use the same key to the password
	if the two hash are the same, func return true and a nil error

	at the end func returne false and a nil error cose password is different than the hash
	*/
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
	/*
	ReadHash function take a stocked hash an return it
	-----------------------------------------------------
	input : a stock hash a type string
	output : the function return the hash key and hash as array of byte and an error if the function fail
	-----------------------------------------------------
	first the func extract the information stocked in the stockhash
	func verifie if the number of information is equal to 4, return an error if not

	func verifie if the curent version of argon is the same used for the stock hash, if the both version arent the same func return an error

	func decode the hash key encode in B64 and see if the length of the key is the same as parametres, return an error if key have diferent length
	than the params
	
	func decode B64 hash and see if is length is the same as params, return an error if not

	after those test the func return the key, hash and a nil err
	*/
	values := strings.Split(StockHash, "$")
	if len(values) != 4 {
		return nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(values[1], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrIncompatibleVersion
	}

	key, err = base64.RawStdEncoding.Strict().DecodeString(values[2])
	if err != nil {
		return nil, nil, err
	}
	if conste.keyLength != uint32(len(key)) {
		return nil, nil, ErrInvalidHash
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[3])
	if err != nil {
		return nil, nil, err
	}
	if conste.hashLength != uint32(len(hash)) {
		return nil, nil, ErrInvalidHash
	}

	return key, hash, nil
}