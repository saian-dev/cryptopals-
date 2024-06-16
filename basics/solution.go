package basics

import (
	"encoding/hex"
	"errors"
	"math/bits"
	"unicode"
)

var letterFrequency = map[rune]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253,
	'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094,
	'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929,
	'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
	'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074, ' ': 0.1,
}

func XOR(msg []byte, key []byte) ([]byte, error) {
	if len(msg) != len(key) {
		return nil, errors.New("lengths should be equal")
	}

	result := make([]byte, 0, len(msg))
	for i := range len(msg) {
		result = append(result, msg[i]^key[i])
	}
	return result, nil
}

func ScoreText(text []byte) float64 {
	if len(text) == 0 {
		return 0
	}

	score := 0.0
	for _, char := range string(text) {
		char := unicode.ToLower(char)
		frequency, ok := letterFrequency[char]
		if ok {
			score += frequency
		}
	}
	return score / float64(len(text))
}

func SingleByteXOR(text []byte) (string, byte) {
	var result_text string
	var key byte
	highestScore := 0.0
	for i := 0; i <= 255; i++ {
		key = byte(i)
		result := make([]byte, 0)
		for i := range len(text) {
			result = append(result, text[i]^key)
		}

		score := ScoreText(result)
		if score > highestScore {
			highestScore = score
			result_text = string(result)
		}
	}

	return result_text, key
}

func RepeatingXOR(msg string, key string) string {
	result := make([]byte, 0)
	byte_key := []byte(key)
	for i, v := range []byte(msg) {
		result = append(result, v^byte_key[i%len(byte_key)])
	}
	return hex.EncodeToString(result)
}

func HummingDistance(a []byte, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("lengths should be equal")
	}

	diff := 0
	for i := range len(a) {
		diff += bits.OnesCount(uint(a[i] ^ b[i]))
	}
	return diff, nil
}
