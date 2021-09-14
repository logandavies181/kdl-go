package main

import (
	"fmt"
	"strings"
)

var test = `
foo bar=r"baz ah" true false null 1 (foobar)"baz"
`

func main() {
	fmt.Println("starting")

	p := NewParser(strings.NewReader(test))
	n, err := p.Parse()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, m := range n {
		fmt.Println(m.Tok, m.Lit)
	}
}
