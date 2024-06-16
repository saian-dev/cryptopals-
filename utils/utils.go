package utils

import (
	"encoding/base64"
	"encoding/hex"
	"os"
)

func ReadFile(fileAt string) []byte {
	data, err := os.ReadFile(fileAt)
	if err != nil {
		panic(err)
	}
	return data
}

func ConvertHexToBase64(hexString string) (string, error) {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func ConvertHexToBytes(hexString string) []byte {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return decoded
}

func ConvertHexToString(src []byte) string {
	return hex.EncodeToString(src)
}
