package main

import (
	"fmt"
	"strings"
)

var test = `
foo bar="baz"
`

func main() {
	fmt.Println("starting")

	p := NewParser(strings.NewReader(test))
	n, err := p.Parse()
	if err != nil {
		fmt.Println("Error:", err)
	}

	if n != nil {
		fmt.Println(n.Id, n.Attr)
	}
}
