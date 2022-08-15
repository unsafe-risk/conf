package lex

import "strconv"

type TokenType uint16

const (
	ILLEGAL TokenType = iota
	EOF
	COMMENT
	COMMENT_LINE

	IDENT      // Identifier
	INT_LIT    // Integer literal
	FLOAT_LIT  // Float literal
	STRING_LIT // String literal
	BOOL_LIT   // Boolean literal
	NULL_LIT   // Null literal (null or nil)

	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LBRACK    // [
	RBRACK    // ]
	COMMA     // ,
	DOT       // .
	COLON     // :
	SEMICOLON // ;
	QUESTION  // ?
	BANG      // !

	// Operators
	PLUS  // +
	MINUS // -
	MUL   // *
	DIV   // /
	MOD   // %

	BIT_AND // &
	BIT_OR  // |
	BIT_XOR // ^
	BIT_NOT // ~
	BIT_LSH // <<
	BIT_RSH // >>

	LT  // <
	GT  // >
	EQ  // ==
	NEQ // !=
	LEQ // <=
	GEQ // >=

	LOGIC_AND // &&
	LOGIC_OR  // ||
	LOGIC_NOT // !
)

var tokens = [...]string{
	ILLEGAL:      "ILLEGAL",
	EOF:          "EOF",
	COMMENT:      "COMMENT",
	COMMENT_LINE: "COMMENT_LINE",
	IDENT:        "IDENT",
	INT_LIT:      "INT_LIT",
	FLOAT_LIT:    "FLOAT_LIT",
	STRING_LIT:   "STRING_LIT",
	BOOL_LIT:     "BOOL_LIT",
	NULL_LIT:     "NULL_LIT",
	LPAREN:       "LPAREN",
	RPAREN:       "RPAREN",
	LBRACE:       "LBRACE",
	RBRACE:       "RBRACE",
	LBRACK:       "LBRACK",
	RBRACK:       "RBRACK",
	COMMA:        "COMMA",
	DOT:          "DOT",
	COLON:        "COLON",
	SEMICOLON:    "SEMICOLON",
	QUESTION:     "QUESTION",
	BANG:         "BANG",
	PLUS:         "PLUS",
	MINUS:        "MINUS",
	MUL:          "MUL",
	DIV:          "DIV",
	MOD:          "MOD",
	BIT_AND:      "BIT_AND",
	BIT_OR:       "BIT_OR",
	BIT_XOR:      "BIT_XOR",
	BIT_NOT:      "BIT_NOT",
	BIT_LSH:      "BIT_LSH",
	BIT_RSH:      "BIT_RSH",
	LT:           "LT",
	GT:           "GT",
	EQ:           "EQ",
	NEQ:          "NEQ",
	LEQ:          "LEQ",
	GEQ:          "GEQ",
	LOGIC_AND:    "LOGIC_AND",
	LOGIC_OR:     "LOGIC_OR",
	LOGIC_NOT:    "LOGIC_NOT",
}

func (t TokenType) String() string {
	if t >= TokenType(len(tokens)) {
		return "token(" + strconv.Itoa(int(t)) + ")"
	}
	return tokens[t]
}

func (t TokenType) IsLiteral() bool {
	return t == INT_LIT || t == FLOAT_LIT || t == STRING_LIT || t == BOOL_LIT || t == NULL_LIT
}

type Token struct {
	Type TokenType
	Pos  Position

	Value string
}

func (t Token) String() string {
	return t.Pos.String() + " | " + t.Type.String() + " " + strconv.Quote(t.Value)
}

func NewToken(pos Position, typ TokenType, value string) Token {
	return Token{
		Type:  typ,
		Pos:   pos,
		Value: value,
	}
}
