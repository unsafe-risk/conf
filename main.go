package main

import (
	"fmt"
	"os"

	"v8.run/go/conf/script/lexer"
	"v8.run/go/conf/script/parser"
	"v8.run/go/conf/script/parser/bsr"
)

var testdata = `
func AddOne(x:u64) u64 {
	return x + 1;
}

let x:u64 = AddOne(1);
if (x == 1) {
	print("AddOne(1) == 1");
} else if (x == 2) {
	print("AddOne(1) == 2");
} else {
	print("AddOne(1) == something else");
}

`

func Walk(b bsr.BSR, fn func(bsr.BSR)) {
	fn(b)
	for _, c := range b.GetAllNTChildren() {
		for _, nt := range c {
			Walk(nt, fn)
		}
	}
}

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

	if p.IsAmbiguous() {
		fmt.Println("ambiguous")
	} else {
		fmt.Println("not ambiguous")
	}
}
