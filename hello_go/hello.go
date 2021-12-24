package main

import (
	"example/hello_go/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
