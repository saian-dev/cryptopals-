package basics

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
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
	return hex.EncodeToString(result), nil
}

func FindMostFrequentSingleByteCharacter(hexString string) (rune, error) {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return 0, err
	}

	var mostFrequent rune
	var maxCount int

	charFrequency := make(map[rune]int)
	for _, char := range string(decoded) {
		if char > 255 {
			continue
		}
		_, ok := charFrequency[char]
		if !ok {
			charFrequency[char] = 1
		} else {
			charFrequency[char] += 1
		}
		if charFrequency[char] > maxCount {
			maxCount = charFrequency[char]
			mostFrequent = char
		}
	}
	return mostFrequent, nil
}

func SingleByteXOR(text []byte, key byte) string {
	result := make([]byte, 0)
	for i := range len(text) {
		result = append(result, text[i]^key)
	}
	return string(result)
}
