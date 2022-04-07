# go-combinations

`go-combinations` is a simple module to create combinations of unsigned integer,
which you can use as indexes on your array (or slice)

## Given M, Select N
- SomeIndexes
  create all possible combinations **except** none.

- AllIndexes
  create all possible combinations where you pick certain number of candidates
  which is actually calling SomeIndexes with all range of selection (1..N)

## Install

```sh
go get github.com/jeongoon/go-combinations

```
## Usage

```go
package main

import "fmt"
import combo "github.com/jeongoon/go-combinations"
/*     ^^^^^ is an alias for this module */

func main() {
        someAlphas := []string{ "A", "B", "C", "D", "E" }
        fmt.Printf( "given elements are:\n%v\n", someAlphas )

        for _, cb := range combo.SomeIndexes(len(someAlphas), 3) {
                fmt.Printf( "[%v %v %v]\n",
                        someAlphas[cb[0]], someAlphas[cb[1]], someAlphas[cb[2]] )
        }
}
```
