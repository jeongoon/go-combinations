# go-combinations

`go-combinations` is a simple module to create combinations of unsigned integer,
which you can use as indexes on your array.

## Given M, Select N
- SomeIndexes
  create all possible combinations **except** none.

- AllIndexes
  create all possible combinations where you pick certain number of candidates
  which is actually calling SomeIndexes with all range of selection (1..N)

## Usage

```go
package main

import "fmt"
import combo "github.com/jeongoon/go-combinations"

func main() {
	fmt.Println(combos.SomeIndexes(7, 5))
}
```
