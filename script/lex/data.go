package lex

import "strconv"

type Source struct {
	Filename string
	RawData  []byte

	Data []rune
}

type Position struct {
	Source *Source
	Line   int
	Column int
}

func (p Position) String() string {
	return p.Source.Filename + ":" + strconv.Itoa(p.Line) + ":" + strconv.Itoa(p.Column)
}

func NewSource(filename string, data []byte) *Source {
	return &Source{
		Filename: filename,
		RawData:  data,
		Data:     []rune(string(data)),
	}
}
