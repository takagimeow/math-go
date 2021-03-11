package main

import (
	"testing"
)

func TestDecToHex(t *testing.T) {
	hex, _ := decToHex(26)
	t.Error(hex)
}
