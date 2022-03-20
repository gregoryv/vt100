package vt100_test

import (
	"fmt"

	"github.com/gregoryv/vt100"
)

func Example() {
	fg := vt100.ForegroundColors()
	bg := vt100.BackgroundColors()
	vt := vt100.Attributes()

	fmt.Println(fg.Red, bg.White, "hello", vt.Reset)
}
