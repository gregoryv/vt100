package vt100

import (
	"testing"
)

func TestSplit(t *testing.T) {
	data := []struct {
		txt string
		exp int
	}{
		{"Hello world", 1},
		{"\033[30mHello", 2},
		{"\033[31mHello", 2},
		{"\033[32mHello", 2},
		{"Hello \033[33m World", 3},
		{"\033[0m Hello \033[35m World", 4},
	}

	for _, d := range data {
		res := Split(d.txt)
		if len(res) != d.exp {
			print("\033[0m")
			t.Errorf("Scan(%q) expected %v, got %v %#v", d.txt, d.exp, len(res), res)
		}

	}
}
