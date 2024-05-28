package basics

import (
	"testing"
)

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := ConvertHexToBase64(input)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if result != expected {
		t.Fatalf("ConvertHexToBase64: %s expected but got %s", expected, result)
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
