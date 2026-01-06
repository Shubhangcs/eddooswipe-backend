package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// PKCS7 padding
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}

// Encrypt AES-128-CBC and return Base64 string
func EncryptAES128CBC(plainText, key, iv string) (string, error) {
	keyBytes := []byte(key)
	ivBytes := []byte(iv)

	if len(keyBytes) != 16 || len(ivBytes) != 16 {
		return "", fmt.Errorf("key and iv must be 16 bytes")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	data := pkcs7Pad([]byte(plainText), block.BlockSize())

	cipherText := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, data)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
