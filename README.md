[vt100](https://pkg.go.dev/github.com/gregoryv/vt100) - Provides vt100 codes and tools

## Quick start

    $ go get github.com/gregoryv/vt100
	
Use color

```go:
import "github.com/gregoryv/vt100"

func main() {
	fmt.Println(vt100.GREEN, "hello", vt100.RESET)
}
```
