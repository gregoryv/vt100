[vt100](https://pkg.go.dev/github.com/gregoryv/vt100) - Provides vt100 codes and tools

## Quick start

    $ go get github.com/gregoryv/vt100
	
Use color

```go:
import "github.com/gregoryv/vt100"

func Example() {
	fg := vt100.ForegroundColors()
	bg := vt100.BackgroundColors()
	vt := vt100.Attributes()

	fmt.Println(fg.Red, bg.White, "hello", vt.Reset)
}
```
