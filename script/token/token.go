
// Package token is generated by GoGLL. Do not edit
package token

import(
    "fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
    typ        Type
    lext, rext int
    input      []rune
}

/*
New returns a new token.
lext is the left extent and rext the right extent of the token in the input.
input is the input slice scanned by the lexer.
*/
func New(t Type, lext, rext int, input []rune) *Token {
    return &Token{
        typ:   t,
        lext:  lext,
        rext:  rext,
        input: input,
    }
}

// GetLineColumn returns the line and column of the left extent of t
func (t *Token) GetLineColumn() (line, col int) {
    line, col = 1, 1
    for j := 0; j < t.lext; j++ {
        switch t.input[j] {
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

// GetInput returns the input from which t was parsed.
func (t *Token) GetInput() []rune {
    return t.input
}

// Lext returns the left extent of t
func (t *Token) Lext() int {
    return t.lext
}

// Literal returns the literal runes of t scanned by the lexer
func (t *Token) Literal() []rune {
    return t.input[t.lext:t.rext]
}

// LiteralString returns string(t.Literal())
func (t *Token) LiteralString() string {
    return string(t.Literal())
}

// LiteralStripEscape returns the literal runes of t scanned by the lexer
func (t *Token) LiteralStripEscape() []rune {
	lit := t.Literal()
	strip := make([]rune, 0, len(lit))
	for i := 0; i < len(lit); i++ {
		if lit[i] == '\\' {
			i++
			switch lit[i] {
			case 't':
				strip = append(strip, '\t')
			case 'r':
				strip = append(strip, '\r')
			case 'n':
				strip = append(strip, '\r')
			default:
				strip = append(strip, lit[i])
			}
		} else {
			strip = append(strip, lit[i])
		}
	}
	return strip
}

// LiteralStringStripEscape returns string(t.LiteralStripEscape())
func (t *Token) LiteralStringStripEscape() string {
	return string(t.LiteralStripEscape())
}

// Rext returns the right extent of t in the input
func (t *Token) Rext() int {
    return t.rext
}

func (t *Token) String() string {
    return fmt.Sprintf("%s (%d,%d) %s",
        t.TypeID(), t.lext, t.rext, t.LiteralString())
}

// Suppress returns true iff t is suppressed by the lexer
func (t *Token) Suppress() bool {
	return Suppress[t.typ]
}

// Type returns the token Type of t
func (t *Token) Type() Type {
    return t.typ
}

// TypeID returns the token Type ID of t. 
// This may be different from the literal of token t.
func (t *Token) TypeID() string {
    return t.Type().ID()
}

// Type is the token type
type Type int

func (t Type) String() string {
    return TypeToString[t]
}

// ID returns the token type ID of token Type t
func (t Type) ID() string {
    return TypeToID[t]
}


const(
    Error  Type = iota  // Error 
    EOF  // $ 
    T_0  // assign 
    T_1  // break 
    T_2  // colon 
    T_3  // comma 
    T_4  // comment 
    T_5  // config 
    T_6  // continue 
    T_7  // dec 
    T_8  // else 
    T_9  // float_literal 
    T_10  // for 
    T_11  // function 
    T_12  // identifier 
    T_13  // if 
    T_14  // in 
    T_15  // inc 
    T_16  // integer_literal 
    T_17  // lbrace 
    T_18  // let 
    T_19  // lparen 
    T_20  // rbrace 
    T_21  // return 
    T_22  // rparen 
    T_23  // semicolon 
    T_24  // string_literal 
    T_25  // while 
    T_26  // whitespace 
)

var TypeToString = []string{ 
    "Error",
    "EOF",
    "T_0",
    "T_1",
    "T_2",
    "T_3",
    "T_4",
    "T_5",
    "T_6",
    "T_7",
    "T_8",
    "T_9",
    "T_10",
    "T_11",
    "T_12",
    "T_13",
    "T_14",
    "T_15",
    "T_16",
    "T_17",
    "T_18",
    "T_19",
    "T_20",
    "T_21",
    "T_22",
    "T_23",
    "T_24",
    "T_25",
    "T_26",
}

var StringToType = map[string] Type { 
    "Error" : Error, 
    "EOF" : EOF, 
    "T_0" : T_0, 
    "T_1" : T_1, 
    "T_2" : T_2, 
    "T_3" : T_3, 
    "T_4" : T_4, 
    "T_5" : T_5, 
    "T_6" : T_6, 
    "T_7" : T_7, 
    "T_8" : T_8, 
    "T_9" : T_9, 
    "T_10" : T_10, 
    "T_11" : T_11, 
    "T_12" : T_12, 
    "T_13" : T_13, 
    "T_14" : T_14, 
    "T_15" : T_15, 
    "T_16" : T_16, 
    "T_17" : T_17, 
    "T_18" : T_18, 
    "T_19" : T_19, 
    "T_20" : T_20, 
    "T_21" : T_21, 
    "T_22" : T_22, 
    "T_23" : T_23, 
    "T_24" : T_24, 
    "T_25" : T_25, 
    "T_26" : T_26, 
}

var TypeToID = []string { 
    "Error", 
    "$", 
    "assign", 
    "break", 
    "colon", 
    "comma", 
    "comment", 
    "config", 
    "continue", 
    "dec", 
    "else", 
    "float_literal", 
    "for", 
    "function", 
    "identifier", 
    "if", 
    "in", 
    "inc", 
    "integer_literal", 
    "lbrace", 
    "let", 
    "lparen", 
    "rbrace", 
    "return", 
    "rparen", 
    "semicolon", 
    "string_literal", 
    "while", 
    "whitespace", 
}

var Suppress = []bool { 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    true, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
}

