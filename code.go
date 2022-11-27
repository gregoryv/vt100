package vt100

import (
	"fmt"
	"strings"
)

func ForegroundColors() ColorCodes {
	return ColorCodes{
		Black:   30,
		Red:     31,
		Green:   32,
		Yellow:  33,
		Blue:    34,
		Magenta: 35,
		Cyan:    36,
		White:   37,
	}
}

func BackgroundColors() ColorCodes {
	return ColorCodes{
		Black:   40,
		Red:     41,
		Green:   42,
		Yellow:  43,
		Blue:    44,
		Magenta: 45,
		Cyan:    46,
		White:   47,
	}
}

type ColorCodes struct {
	Black   Code
	Red     Code
	Green   Code
	Yellow  Code
	Blue    Code
	Magenta Code
	Cyan    Code
	White   Code
}

func (c *ColorCodes) ByName(v string) Code {
	m := map[string]Code{
		"black":   c.Black,
		"red":     c.Red,
		"green":   c.Green,
		"yellow":  c.Yellow,
		"blue":    c.Blue,
		"magenta": c.Magenta,
		"cyan":    c.Cyan,
		"white":   c.White,
	}
	return m[strings.ToLower(v)]
}

func Attributes() AttributeCodes {
	return AttributeCodes{
		Reset:      0,
		Bright:     1,
		Dim:        2,
		Underscore: 4,
		Blink:      5,
		Reverse:    7,
		Hidden:     8,
	}
}

type AttributeCodes struct {
	Reset      Code
	Bright     Code
	Dim        Code
	Underscore Code
	Blink      Code
	Reverse    Code
	Hidden     Code
}

type Code uint

func (me Code) Bytes() []byte {
	return []byte(me.String())
}

func (me Code) String() string {
	return fmt.Sprintf("\033[%vm", uint(me))
}
