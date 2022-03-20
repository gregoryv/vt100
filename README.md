[vt100](https://pkg.go.dev/github.com/gregoryv/vt100) - Provides vt100 codes and tools

## Quick start

    $ go get github.com/gregoryv/vt100/code
	
Use color

```go:
import "github.com/gregoryv/vt100"

func main() {
	fmt.Println(Green, "hello", Reset)
}
```
