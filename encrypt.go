package encrypt

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"
)

func HASH(content string, hashType string, isHex bool) (res string, err error) {
	var (
		hashInstance hash.Hash
		arr          []byte
	)
	switch strings.ToLower(hashType) {
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "ripemd160":
		hashInstance = crypto.RIPEMD160.New()
	case "sha256":
		hashInstance = crypto.SHA256.New()
	case "sha512":
		hashInstance = crypto.SHA512.New()
	default:
		err = errors.New("not support crypto")
		return
	}
	if isHex {
		if arr, err = hex.DecodeString(content); err != nil {
			return
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(content))
	}
	res = fmt.Sprintf("%x", hashInstance.Sum(nil))
	return
}

func HASHDouble(content string, hashType string, isHex bool) (res string, err error) {
	var (
		hashInstance hash.Hash
		arr          []byte
		bytes        []byte
	)
	switch strings.ToLower(hashType) {
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "ripemd160":
		hashInstance = crypto.RIPEMD160.New()
	case "sha256":
		hashInstance = crypto.SHA256.New()
	case "sha512":
		hashInstance = crypto.SHA512.New()
	default:
		err = errors.New("not support crypto")
		return
	}
	if isHex {
		if arr, err = hex.DecodeString(content); err != nil {
			return
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(content))
	}
	bytes = hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(bytes)
	bytes = hashInstance.Sum(nil)

	res = fmt.Sprintf("%x", bytes)
	return
}
