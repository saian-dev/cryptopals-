package basics

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func ConvertHexToBase64(hexString string) (string, error) {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func XOR(msg string, key string) (string, error) {
	if len(msg) != len(key) {
		return "", errors.New("lengths should be equal")
	}

	decoded_msg, err := hex.DecodeString(msg)
	if err != nil {
		return "", err
	}

	decoded_keys, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	result := make([]byte, 0, len(decoded_msg))
	for i := range len(decoded_msg) {
		result = append(result, decoded_msg[i]^decoded_keys[i])
	}
	fmt.Println(string(result))
	return hex.EncodeToString(result), nil
}
