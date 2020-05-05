// aes.go
package ssui

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return nil
	}
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func Encrypt(pass, key string) (string, error) {
	var aeskey = []byte(key)
	p := []byte(pass)
	xpass, err := AesEncrypt(p, aeskey)
	if err != nil {
		return "", err
	}

	pass64 := base64.URLEncoding.EncodeToString(xpass)
	return pass64, nil
}

func Decrypt(pass64, key string) (string, error) {
	var aeskey = []byte(key)
	bytesPass, err := base64.URLEncoding.DecodeString(pass64)
	if err != nil {
		return "", err
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		return "", err
	}
	return string(tpass), nil
}
