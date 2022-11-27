package vt100

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	c := Code(31)
	s := c.String()
	b := c.Bytes()
	if string(b) != s {
		t.Error("String and Bytes don't match")
	}
}

func TestForegroundColors(t *testing.T) {
	fg := ForegroundColors()
	bg := BackgroundColors()

	if fg.Black == bg.Black {
		t.Error("foreground and background codes are the same")
	}
}

func TestVisualInspectFG(t *testing.T) {
	defer func() {
		vt := Attributes()
		fmt.Println(vt.Reset)
	}()

	fg := ForegroundColors()
	t.Log("Black", fg.ByName("Black"), "Black")
	t.Log("Red", fg.ByName("Red"), "Red")
	t.Log("Green", fg.ByName("Green"), "Green")
	t.Log("Yellow", fg.ByName("Yellow"), "Yellow")
	t.Log("Blue", fg.ByName("Blue"), "Blue")
	t.Log("Magenta", fg.ByName("Magenta"), "Magenta")
	t.Log("Cyan", fg.ByName("Cyan"), "Cyan")
	t.Log("White", fg.ByName("White"), "White")
	// t.Fail()
}

func TestVisualInspectBG(t *testing.T) {
	defer func() {
		vt := Attributes()
		t.Log(vt.Reset)
	}()

	bg := BackgroundColors()
	t.Log(bg.Black)
	t.Log(bg.Red)
	t.Log(bg.Green)
	t.Log(bg.Yellow)
	t.Log(bg.Blue)
	t.Log(bg.Magenta)
	t.Log(bg.Cyan)
	t.Log(bg.White)
	// t.Fail()
}

func BenchmarkCode(b *testing.B) {
	c := Code(31)
	for i := 0; i < b.N; i++ {
		c.Bytes()
	}
}
