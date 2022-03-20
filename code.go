package vt100

import "fmt"

type Code uint

func (me Code) Bytes() []byte {
	return []byte(me.String())
}

func (me Code) String() string {
	return fmt.Sprintf("\033[%vm", uint(me))
}

func Background(color Code) Code {
	return Code(color + 10)
}

const (
	CodeReset      Code = 0 //	Reset all attributes
	CodeBright     Code = 1 //	Bright
	CodeDim        Code = 2 //	Dim
	CodeUnderscore Code = 4 //	Underscore
	CodeBlink      Code = 5 //	Blink
	CodeReverse    Code = 7 //	Reverse
	CodeHidden     Code = 8 //	Hidden

	// Foreground Colours
	Black   Code = 30 //	Black
	Red     Code = 31 //	Red
	Green   Code = 32 //	Green
	Yellow  Code = 33 //	Yellow
	Blue    Code = 34 //	Blue
	Magenta Code = 35 //	Magenta
	Cyan    Code = 36 //	Cyan
	White   Code = 37 //	White
)

/*
https://www2.ccs.neu.edu/research/gpc/VonaUtils/vona/terminal/vtansi.htm#colors
*/
