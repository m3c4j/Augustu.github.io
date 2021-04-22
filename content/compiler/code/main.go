package main

import "fmt"

func main() {
	// test := "1 + 2 + 3"
	// test := "1+2+3"

	// p := &Parser{}
	// p.Parser(test)
	// err := p.expr()
	// fmt.Println(err)

	test := "a + 12 = true"

	l := NewLexer(test)
	for t, fin := l.scan(); !fin; t, fin = l.scan() {
		fmt.Println(t.String())
		fmt.Printf("%#v\n", t)
	}
}
