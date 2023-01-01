package parser

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	Text        string
	position    int
	diagnostics []string
}

func (lex Lexer) current() string {
	if lex.position >= len(lex.Text) {
		return `\0`
	}
	return string(lex.Text[lex.position])
}

func (lex *Lexer) next() {
	lex.position++
}

// NextToken : reads in the next token needed
func (lex *Lexer) NextToken() SyntaxToken {
	// + - * /
	// numbers
	// <whitespace>
	if lex.position >= len(lex.Text) {
		return SyntaxToken{
			Kind_:    EndOfFileToken,
			position: lex.position,
			Text:     `\0`,
		}
	}
	start := lex.position
	if _, err := strconv.Atoi(lex.current()); err == nil {
		for {
			_, err := strconv.Atoi(lex.current())
			if err == nil {
				lex.next()
			} else {
				break
			}
		}

		text := lex.Text[start:lex.position]
		value, err := strconv.Atoi(text)
		if err != nil {
			lex.diagnostics = append(lex.diagnostics, fmt.Sprintf("The number %s is not a valid int", text))
		}
		return SyntaxToken{
			Kind_:    NumberToken,
			position: start,
			Text:     text,
			Value:    value,
		}
	}
	if lex.current() == " " {
		for {
			if lex.current() == " " {
				lex.next()
			} else {
				break
			}
		}

		text := lex.Text[start:lex.position]
		return SyntaxToken{
			Kind_:    WhitespaceToken,
			position: start,
			Text:     text,
		}
	} else if lex.current() == "+" {
		lex.next()
		return SyntaxToken{
			Kind_:    PlusToken,
			position: lex.position - 1,
			Text:     "+",
		}
	} else if lex.current() == "-" {
		lex.next()
		return SyntaxToken{
			Kind_:    MinusToken,
			position: lex.position - 1,
			Text:     "-",
		}
	} else if lex.current() == "*" {
		lex.next()
		return SyntaxToken{
			Kind_:    StarToken,
			position: lex.position - 1,
			Text:     "*",
		}
	} else if lex.current() == "/" {
		lex.next()
		return SyntaxToken{
			Kind_:    ForwardSlashToken,
			position: lex.position - 1,
			Text:     "/",
		}
	} else if lex.current() == "(" {
		lex.next()
		return SyntaxToken{
			Kind_:    OpenParenthesisToken,
			position: lex.position - 1,
			Text:     "(",
		}
	} else if lex.current() == ")" {
		lex.next()
		return SyntaxToken{
			Kind_:    CloseParenthesisToken,
			position: lex.position - 1,
			Text:     ")",
		}
	}
	lex.diagnostics = append(lex.diagnostics, fmt.Sprintf("ERROR: bad character input: %s", lex.current()))
	lex.next()
	return SyntaxToken{
		Kind_:    BadToken,
		position: lex.position - 1,
		Text:     string(lex.Text[lex.position-1]),
	}
}