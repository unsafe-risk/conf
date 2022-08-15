
// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"v8.run/go/conf/script/token"
)

type state int

const nullState state = -1


// Lexer contains both the input slice of runes and the slice of tokens
// parsed from the input
type Lexer struct {
	// I is the input slice of runes
	I      []rune

	// Tokens is the slice of tokens constructed by the lexer from I
	Tokens []*token.Token
}

/*
NewFile constructs a Lexer created from the input file, fname. 

If the input file is a markdown file NewFile process treats all text outside
code blocks as whitespace. All text inside code blocks are treated as input text.

If the input file is a normal text file NewFile treats all text in the inputfile
as input text.
*/
func NewFile(fname string) *Lexer {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	input := []rune(string(buf))
	if strings.HasSuffix(fname, ".md") {
		loadMd(input)
	}
	return New(input)
}

func loadMd(input []rune) {
	i := 0
	text := true
	for i < len(input) {
		if i <= len(input)-3 && input[i] == '`' && input[i+1] == '`' && input[i+2] == '`' {
			text = !text
			for j := 0; j < 3; j++ {
				input[i+j] = ' '
			}
			i += 3
		}
		if i < len(input) {
			if text {
				if input[i] == '\n' {
					input[i] = '\n'
				} else {
					input[i] = ' '
				}
			}
			i += 1
		}
	}
}

/*
New constructs a Lexer from a slice of runes. 

All contents of the input slice are treated as input text.
*/
func New(input []rune) *Lexer {
	lex := &Lexer{
		I:      input,
		Tokens: make([]*token.Token, 0, 2048),
	}
	lext := 0
	for lext < len(lex.I) {
		for lext < len(lex.I) && unicode.IsSpace(lex.I[lext]) {
			lext++
		}
		if lext < len(lex.I) {
			tok := lex.scan(lext)
			lext = tok.Rext()
			if !tok.Suppress() {
				lex.addToken(tok)
			}
		}
	}
	lex.add(token.EOF, len(input), len(input))
	return lex
}

func (l *Lexer) scan(i int) *token.Token {
	// fmt.Printf("lexer.scan(%d)\n", i)
	s, typ, rext := nullState, token.Error, i+1
	if i < len(l.I) {
		// fmt.Printf("  rext %d, i %d\n", rext, i)
		s = nextState[0](l.I[i])
	}
	for s != nullState {
		if rext >= len(l.I) {
			typ = accept[s]
			s = nullState
		} else {
			typ = accept[s]
			s = nextState[s](l.I[rext])
			if s != nullState || typ == token.Error {
				rext++
			}
		}
	}
	tok := token.New(typ, i, rext, l.I)
	// fmt.Printf("  %s\n", tok)
	return tok
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetLineColumnOfToken returns the line and column of token[i] in the imput
func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext())
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext():l.Tokens[rext].Rext()])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.New(t, lext, rext, l.I))
}

func (l *Lexer) addToken(tok *token.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = []token.Type{ 
	token.Error, 
	token.Error, 
	token.T_3, 
	token.T_18, 
	token.T_21, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_15, 
	token.T_2, 
	token.T_22, 
	token.T_0, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_16, 
	token.T_19, 
	token.T_25, 
	token.T_11, 
	token.Error, 
	token.T_23, 
	token.Error, 
	token.T_14, 
	token.T_15, 
	token.T_6, 
	token.T_8, 
	token.Error, 
	token.T_3, 
	token.T_8, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_15, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_10, 
	token.Error, 
	token.T_12, 
	token.T_13, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_8, 
	token.T_15, 
	token.T_15, 
	token.T_15, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_9, 
	token.T_17, 
	token.Error, 
	token.Error, 
	token.T_3, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_7, 
	token.Error, 
	token.Error, 
	token.T_1, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.T_24, 
	token.T_4, 
	token.Error, 
	token.T_20, 
	token.Error, 
	token.T_5, 
}

var nextState = []func(r rune) state{ 
	// Set0
	func(r rune) state {
		switch { 
		case r == '"':
			return 1 
		case r == '#':
			return 2 
		case r == '(':
			return 3 
		case r == ')':
			return 4 
		case r == '+':
			return 5 
		case r == '-':
			return 6 
		case r == '.':
			return 7 
		case r == '/':
			return 8 
		case r == '0':
			return 9 
		case r == ':':
			return 10 
		case r == ';':
			return 11 
		case r == '=':
			return 12 
		case r == 'b':
			return 13 
		case r == 'c':
			return 14 
		case r == 'e':
			return 15 
		case r == 'f':
			return 16 
		case r == 'i':
			return 17 
		case r == 'l':
			return 18 
		case r == 'r':
			return 19 
		case r == 'w':
			return 20 
		case r == '{':
			return 21 
		case r == '}':
			return 22 
		case any(r, []rune{'\t','\n','\r',' '}):
			return 23 
		case unicode.IsUpper(r):
			return 24 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 9 
		case any(r, []rune{'+','-'}):
			return 25 
		}
		return nullState
	}, 
	// Set1
	func(r rune) state {
		switch { 
		case r == '"':
			return 26 
		case r == '\\':
			return 27 
		case not(r, []rune{'"','\\'}):
			return 1 
		}
		return nullState
	}, 
	// Set2
	func(r rune) state {
		switch { 
		case not(r, []rune{'\n'}):
			return 2 
		}
		return nullState
	}, 
	// Set3
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set4
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set5
	func(r rune) state {
		switch { 
		case r == '+':
			return 28 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 29 
		}
		return nullState
	}, 
	// Set6
	func(r rune) state {
		switch { 
		case r == '-':
			return 30 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 29 
		}
		return nullState
	}, 
	// Set7
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 31 
		}
		return nullState
	}, 
	// Set8
	func(r rune) state {
		switch { 
		case r == '*':
			return 32 
		case r == '/':
			return 33 
		}
		return nullState
	}, 
	// Set9
	func(r rune) state {
		switch { 
		case r == '.':
			return 34 
		case r == 'b':
			return 35 
		case r == 'o':
			return 36 
		case r == 'x':
			return 37 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9','_'}):
			return 38 
		}
		return nullState
	}, 
	// Set10
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set11
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set12
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set13
	func(r rune) state {
		switch { 
		case r == 'r':
			return 39 
		}
		return nullState
	}, 
	// Set14
	func(r rune) state {
		switch { 
		case r == 'o':
			return 40 
		}
		return nullState
	}, 
	// Set15
	func(r rune) state {
		switch { 
		case r == 'l':
			return 41 
		}
		return nullState
	}, 
	// Set16
	func(r rune) state {
		switch { 
		case r == 'n':
			return 42 
		case r == 'o':
			return 43 
		}
		return nullState
	}, 
	// Set17
	func(r rune) state {
		switch { 
		case r == 'f':
			return 44 
		case r == 'n':
			return 45 
		}
		return nullState
	}, 
	// Set18
	func(r rune) state {
		switch { 
		case r == 'e':
			return 46 
		}
		return nullState
	}, 
	// Set19
	func(r rune) state {
		switch { 
		case r == 'e':
			return 47 
		}
		return nullState
	}, 
	// Set20
	func(r rune) state {
		switch { 
		case r == 'h':
			return 48 
		}
		return nullState
	}, 
	// Set21
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set22
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set23
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set24
	func(r rune) state {
		switch { 
		case r == '_':
			return 24 
		case unicode.IsLetter(r):
			return 24 
		case unicode.IsNumber(r):
			return 24 
		}
		return nullState
	}, 
	// Set25
	func(r rune) state {
		switch { 
		case r == '+':
			return 28 
		case r == '-':
			return 30 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 29 
		}
		return nullState
	}, 
	// Set26
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set27
	func(r rune) state {
		switch { 
		case true:
			return 1 
		}
		return nullState
	}, 
	// Set28
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set29
	func(r rune) state {
		switch { 
		case r == '.':
			return 34 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9','_'}):
			return 38 
		}
		return nullState
	}, 
	// Set30
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set31
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 31 
		}
		return nullState
	}, 
	// Set32
	func(r rune) state {
		switch { 
		case r == '*':
			return 49 
		case not(r, []rune{'*'}):
			return 32 
		}
		return nullState
	}, 
	// Set33
	func(r rune) state {
		switch { 
		case not(r, []rune{'\n'}):
			return 33 
		}
		return nullState
	}, 
	// Set34
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9'}):
			return 50 
		}
		return nullState
	}, 
	// Set35
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1'}):
			return 51 
		}
		return nullState
	}, 
	// Set36
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7'}):
			return 52 
		}
		return nullState
	}, 
	// Set37
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','a','b','c','d','e','f'}):
			return 53 
		}
		return nullState
	}, 
	// Set38
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9','_'}):
			return 38 
		}
		return nullState
	}, 
	// Set39
	func(r rune) state {
		switch { 
		case r == 'e':
			return 54 
		}
		return nullState
	}, 
	// Set40
	func(r rune) state {
		switch { 
		case r == 'n':
			return 55 
		}
		return nullState
	}, 
	// Set41
	func(r rune) state {
		switch { 
		case r == 's':
			return 56 
		}
		return nullState
	}, 
	// Set42
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set43
	func(r rune) state {
		switch { 
		case r == 'r':
			return 57 
		}
		return nullState
	}, 
	// Set44
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set45
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set46
	func(r rune) state {
		switch { 
		case r == 't':
			return 58 
		}
		return nullState
	}, 
	// Set47
	func(r rune) state {
		switch { 
		case r == 't':
			return 59 
		}
		return nullState
	}, 
	// Set48
	func(r rune) state {
		switch { 
		case r == 'i':
			return 60 
		}
		return nullState
	}, 
	// Set49
	func(r rune) state {
		switch { 
		case r == '/':
			return 61 
		case not(r, []rune{'/'}):
			return 32 
		}
		return nullState
	}, 
	// Set50
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set51
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','_'}):
			return 51 
		}
		return nullState
	}, 
	// Set52
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','_'}):
			return 52 
		}
		return nullState
	}, 
	// Set53
	func(r rune) state {
		switch { 
		case any(r, []rune{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','_','a','b','c','d','e','f'}):
			return 53 
		}
		return nullState
	}, 
	// Set54
	func(r rune) state {
		switch { 
		case r == 'a':
			return 62 
		}
		return nullState
	}, 
	// Set55
	func(r rune) state {
		switch { 
		case r == 'f':
			return 63 
		case r == 't':
			return 64 
		}
		return nullState
	}, 
	// Set56
	func(r rune) state {
		switch { 
		case r == 'e':
			return 65 
		}
		return nullState
	}, 
	// Set57
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set58
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set59
	func(r rune) state {
		switch { 
		case r == 'u':
			return 66 
		}
		return nullState
	}, 
	// Set60
	func(r rune) state {
		switch { 
		case r == 'l':
			return 67 
		}
		return nullState
	}, 
	// Set61
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set62
	func(r rune) state {
		switch { 
		case r == 'k':
			return 68 
		}
		return nullState
	}, 
	// Set63
	func(r rune) state {
		switch { 
		case r == 'i':
			return 69 
		}
		return nullState
	}, 
	// Set64
	func(r rune) state {
		switch { 
		case r == 'i':
			return 70 
		}
		return nullState
	}, 
	// Set65
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set66
	func(r rune) state {
		switch { 
		case r == 'r':
			return 71 
		}
		return nullState
	}, 
	// Set67
	func(r rune) state {
		switch { 
		case r == 'e':
			return 72 
		}
		return nullState
	}, 
	// Set68
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set69
	func(r rune) state {
		switch { 
		case r == 'g':
			return 73 
		}
		return nullState
	}, 
	// Set70
	func(r rune) state {
		switch { 
		case r == 'n':
			return 74 
		}
		return nullState
	}, 
	// Set71
	func(r rune) state {
		switch { 
		case r == 'n':
			return 75 
		}
		return nullState
	}, 
	// Set72
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set73
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set74
	func(r rune) state {
		switch { 
		case r == 'u':
			return 76 
		}
		return nullState
	}, 
	// Set75
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set76
	func(r rune) state {
		switch { 
		case r == 'e':
			return 77 
		}
		return nullState
	}, 
	// Set77
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
}
