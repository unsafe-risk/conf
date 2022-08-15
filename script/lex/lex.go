package lex

type Lexer struct {
	position Position
	last     Token

	absPos int
}

func NewLexer(s *Source) *Lexer {
	return &Lexer{
		position: Position{
			Source: s,
			Line:   1,
			Column: 0,
		},
		absPos: 0,
	}
}

func (l *Lexer) getChar() rune {
	if l.absPos >= len(l.position.Source.Data) {
		return -1
	}
	r := l.position.Source.Data[l.absPos]
	l.absPos++
	if r == '\n' {
		l.position.Line++
		l.position.Column = 1
	} else {
		l.position.Column++
	}
	return r
}

func (l *Lexer) peekChar() rune {
	if l.absPos >= len(l.position.Source.Data) {
		return -1
	}
	return l.position.Source.Data[l.absPos]
}

func (l *Lexer) Next() Token {
	char := l.getChar()
	if char == -1 {
		return NewToken(l.position, EOF, "")
	}

	var t Token = NewToken(l.position, ILLEGAL, "")
	switch char {
	case '/':
		// COMMENT or COMMENT_LINE or MUL
		switch l.peekChar() {
		case '/':
			// COMMENT_LINE
			pos := l.position
			start := l.absPos
			for {
				char = l.getChar()
				if char == -1 || char == '\n' {
					break
				}
			}
			t = NewToken(pos, COMMENT_LINE, string(l.position.Source.Data[start:l.absPos]))
		case '*':
			// COMMENT
			pos := l.position
			start := l.absPos
			for {
				char = l.getChar()
				if char == -1 {
					t = NewToken(pos, ILLEGAL, "unexpected EOF")
					return t
				}
				if char == '*' && l.peekChar() == '/' {
					l.getChar()
					break
				}
			}

			t = NewToken(pos, COMMENT, string(l.position.Source.Data[start:l.absPos]))
		default:
			// MUL
			t = NewToken(l.position, MUL, "*")
		}

	case '"':
		// STRING_LIT
		pos := l.position
		start := l.absPos
		var isEscaped bool
		for {
			char = l.getChar()
			if char == -1 {
				t = NewToken(pos, ILLEGAL, "unexpected EOF")
				return t
			}
			if char == '"' && !isEscaped {
				break
			}
			if char == '\\' && !isEscaped {
				isEscaped = true
			} else {
				isEscaped = false
			}
		}
		t = NewToken(pos, STRING_LIT, string(l.position.Source.Data[start:l.absPos]))
	}

	return t
}
