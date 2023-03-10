package syntax

import (
	"reflect"
	"strconv"
	"unicode"

	"github.com/deven96/welt/diagnostic"
)

type lexer struct {
	Text        string
	position    int
	diagnostics diagnostic.DiagnosticsBag
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (lex lexer) isUnescapedQuote() bool {
	if lex.current() == `"` && lex.lookbehind() != `\` {
		return true
	}
	return false
}

func (lex lexer) lookbehind() string {
	return lex.peek(-1)
}

func (lex lexer) current() string {
	return lex.peek(0)
}

func (lex *lexer) next() {
	lex.position++
}

func (lex lexer) peek(offset int) string {
	index := lex.position + offset
	if index >= len(lex.Text) {
		return EndOfFileChar
	}
	if index < 0 {
		index = 0
	}
	return string(lex.Text[index])
}

// Lex : reads in the next token needed
func (lex *lexer) Lex() SyntaxToken {
	// + - * /
	// numbers
	// <whitespace>
	if lex.position >= len(lex.Text) {
		return SyntaxToken{
			Kind_:    EndOfFileToken,
			position: lex.position - 1,
			Text:     EndOfFileChar,
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
			var a int
			lex.diagnostics.ReportInvalidNumber(diagnostic.TextSpan{Start: start, Length: lex.position - start}, text, reflect.TypeOf(a))
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
	}

	if IsLetter(lex.current()) {
		for {
			if IsLetter(lex.current()) {
				lex.next()
			} else {
				break
			}
		}

		text := lex.Text[start:lex.position]
		return SyntaxToken{
			Kind_:    KeyWordRecognition(text),
			position: start,
			Text:     text,
		}
	}

	if lex.isUnescapedQuote() {
		lex.next()
		// Is only valid quoted expression when not escaped
		completed := false
		for {
			// Reached end of file or completed
			if lex.isUnescapedQuote() {
				completed = true
			}
			if lex.position == (len(lex.Text)-1) || completed {
				break
			}
			lex.next()
		}
		if completed {
			lex.next()
			quotedText := lex.Text[start:lex.position]
			return SyntaxToken{
				Kind_:    QuotedIdentifierToken,
				position: start,
				Text:     quotedText,
			}
		} else {
			text := lex.Text[start : lex.position+1]
			lex.diagnostics.ReportEOLStringLiteral(diagnostic.TextSpan{Start: start, Length: lex.position - start + 1}, text)
		}
	}

	current := lex.current()
	lex.next()

	switch current {
	case "+":
		return SyntaxToken{
			Kind_:    PlusToken,
			position: lex.position - 1,
			Text:     "+",
		}
	case "-":
		return SyntaxToken{
			Kind_:    MinusToken,
			position: lex.position - 1,
			Text:     "-",
		}
	case "*":
		return SyntaxToken{
			Kind_:    StarToken,
			position: lex.position - 1,
			Text:     "*",
		}
	case "/":
		return SyntaxToken{
			Kind_:    ForwardSlashToken,
			position: lex.position - 1,
			Text:     "/",
		}
	case "%":
		return SyntaxToken{
			Kind_:    ModuloToken,
			position: lex.position - 1,
			Text:     "%",
		}
	case "(":
		return SyntaxToken{
			Kind_:    OpenParenthesisToken,
			position: lex.position - 1,
			Text:     "(",
		}
	case ")":
		return SyntaxToken{
			Kind_:    CloseParenthesisToken,
			position: lex.position - 1,
			Text:     "(",
		}
	case "!":
		if lex.current() == "=" {
			// move token once more to the next
			lex.next()
			return SyntaxToken{
				Kind_:    BangEqualToken,
				position: lex.position - 2,
				Text:     "!=",
			}
		}
		return SyntaxToken{
			Kind_:    BangToken,
			position: lex.position - 1,
			Text:     "!",
		}
	case "&":
		if lex.current() == "&" {
			// move token once more to the next
			lex.next()
			return SyntaxToken{
				Kind_:    DoubleAmpersandToken,
				position: lex.position - 2,
				Text:     "&&",
			}
		}
		fallthrough
	case "|":
		if lex.current() == "|" {
			// move token once more to the next
			lex.next()
			return SyntaxToken{
				Kind_:    DoublePipeToken,
				position: lex.position - 2,
				Text:     "||",
			}
		}
		fallthrough
	case "=":
		if lex.current() == "=" {
			// move token once more to the next
			lex.next()
			return SyntaxToken{
				Kind_:    DoubleEqualToken,
				position: lex.position - 2,
				Text:     "==",
			}
		}
		return SyntaxToken{
			Kind_:    EqualsToken,
			position: lex.position - 1,
			Text:     "=",
		}
	default:
		character := lex.Text[lex.position-1]
		lex.diagnostics.ReportBadCharacter(lex.position-1, character)
		return SyntaxToken{
			Kind_:    BadToken,
			position: lex.position - 1,
			Text:     string(character),
		}
	}

}
