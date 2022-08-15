package main

import (
	"fmt"
	"os"

	"v8.run/go/conf/script/lexer"
	"v8.run/go/conf/script/parser"
)

var testdata = `
{
	for (;;) {
		let X:U64 = 1;
		let config VAR_0:U64 = "Hello World";
	}
}
`

func main() {
	l := lexer.New([]rune(testdata))
	for _, tok := range l.Tokens {
		fmt.Printf("%s\n", tok)
	}
	p, errs := parser.Parse(l)
	for _, err := range errs {
		fmt.Println("error:", err)
	}
	if len(errs) > 0 {
		os.Exit(1)
		return
	}
	for _, stmt := range p.GetRoots() {
		fmt.Println(stmt)
	}
}
