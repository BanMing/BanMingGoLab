package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))
}
