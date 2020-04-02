package main

import (
	"testing"
)

func TestSeedGen(t *testing.T) {
	expected := 24
	size := len(seedGen(16))

	if size <= 23 {
		t.Error("Test failed, expected size", expected, "but got", size)
	}
}
