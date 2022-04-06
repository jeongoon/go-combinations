# go-combinations

go-combinations is a simple module to create combinations of unsigned integer,
which you can use as indexes on your array.

- all_combinations_index
  create all possible combinations **except** none.

- combinations_index
  create all possible combinations where you pick certain number of candidates

## Usage

```go
package main

import "fmt"
import combo "github.com/jeongoon/go-combinations"

func main() {
	fmt.Println(combos.combinations_index(7, 5))
}
```
