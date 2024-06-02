package basics

import (
	"encoding/hex"
	"testing"
	"unicode"
)

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := ConvertHexToBase64(input)
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
	result, err := XOR(msg, key)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if result != expected {
		t.Fatalf("XOR: %s expected but got %s", expected, result)
	}
}

func TestSingleByteXOR(t *testing.T) {
	expected := "Cooking MC's like a pound of bacon"

	msg := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	msg_bytes, err := hex.DecodeString(msg)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	key, _ := FindMostFrequentSingleByteCharacter(msg)
	// "x" gives the meaningful text closest to original
	// "X" returns original hence byte(unicode.ToUpper(key))
	// could probably just brute force
	result := SingleByteXOR(msg_bytes, byte(unicode.ToUpper(key)))
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

func TestHummingDistance(t *testing.T) {
	expected := 37

	a := "this is a test"
	b := "wokka wokka!!!"
	result, _ := HummingDistance(a, b)
	if result != expected {
		t.Fatalf("XOR: '%v' expected but got '%v'", expected, result)
	}
}
