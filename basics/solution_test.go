package basics

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/sayan-feb27/cryptopals/utils"
)

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := utils.ConvertHexToBase64(input)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != expected {
		t.Fatalf("ConvertHexToBase64: %v expected but got %v", expected, result)
	}
}

func TestXOR(t *testing.T) {
	msg, key := "1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	msg_as_bytes, key_as_bytes := utils.ConvertHexToBytes(msg), utils.ConvertHexToBytes(key)
	result, err := XOR(msg_as_bytes, key_as_bytes)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if utils.ConvertHexToString(result) != expected {
		t.Fatalf("XOR: %s expected but got %s", expected, result)
	}
}

func TestSingleByteXOR(t *testing.T) {
	expected := "Cooking MC's like a pound of bacon"

	msg := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	msg_bytes := utils.ConvertHexToBytes(msg)

	result, _ := SingleByteXOR(msg_bytes)
	if result != expected {
		t.Fatalf("XOR: '%v' expected but got '%v'", expected, result)
	}
}

func TestRepeatingXOR(t *testing.T) {
	expected := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272` +
		`a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	msg := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	result := RepeatingXOR(msg, key)
	if result != expected {
		t.Fatalf("XOR: '%v' expected but got '%v'", expected, result)
	}
}

func TestDetectSingleCharacterXOR(t *testing.T) {
	expected := "Now that the party is jumping"
	file_path, _ := filepath.Abs("./challenge_4_data.txt")
	file_data := utils.ReadFile(file_path)

	result := ""
	highestScore := 0.0
	text := strings.Split(string(file_data), "\n")
	for _, line := range text {
		line_as_bytes := utils.ConvertHexToBytes(line[0 : len(line)-2])
		decoded_line, _ := SingleByteXOR(line_as_bytes)

		score := ScoreText([]byte(decoded_line))
		if score > highestScore {
			highestScore = score
			result = decoded_line
		}
	}

	if result != expected {
		t.Fatalf("XOR: '%v' expected but got '%v'", expected, result)
	}
}

func TestHummingDistance(t *testing.T) {
	expected := 37

	a := "this is a test"
	b := "wokka wokka!!!"
	result, _ := HummingDistance([]byte(a), []byte(b))
	if result != expected {
		t.Fatalf("XOR: '%v' expected but got '%v'", expected, result)
	}
}
