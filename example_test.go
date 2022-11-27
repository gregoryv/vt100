package vt100_test

import (
	"fmt"
	"os"

	"github.com/gregoryv/vt100"
)

func Example() {
	fg := vt100.ForegroundColors()
	bg := vt100.BackgroundColors()
	vt := vt100.Attributes()

	fmt.Println(fg.Red, bg.White, "hello", vt.Reset)
}

func ExampleParseCodeBytes() {
	attr, _ := vt100.ParseCodeBytes(".*:red;blink;underscore;bgyellow")
	os.Stdout.Write(attr)
	vt := vt100.Attributes()
	fmt.Println("hello", vt.Reset)
}
