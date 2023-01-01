package parser

import "fmt"

type Parser struct {
	Tokens      []SyntaxToken
	position    int
	diagnostics []string
}

func NewParser(text string) Parser {
	p := Parser{}
	lexer := Lexer{Text: text}
	tokens := []SyntaxToken{}
	for {
		token := lexer.NextToken()
		if token.Kind_ != WhitespaceToken && token.Kind_ != BadToken {
			tokens = append(tokens, token)
		}
		if token.Kind_ == EndOfFileToken {
			break
		}
	}
	p.Tokens = tokens
	p.diagnostics = append(p.diagnostics, lexer.diagnostics...)
	return p
}

func (p *Parser) NextToken() SyntaxToken {
	current := p.Current()
	p.position++
	return current
}

func (p *Parser) match(kind SyntaxKind) SyntaxToken {
	if p.Current().Kind_ == kind {
		return p.NextToken()
	}
	p.diagnostics = append(p.diagnostics, fmt.Sprintf("ERROR: Unexpected token <%s>, expected %s", p.Current().Kind(), kind))
	return SyntaxToken{
		Kind_:    kind,
		position: p.Current().position,
	}
}

func (p *Parser) parsePrimaryExpression() ExpressionSyntax {
	if p.Current().Kind() == OpenParenthesisToken {
		left := p.NextToken()
		expression := p.parseTerm()
		right := p.match(CloseParenthesisToken)
		return ParenthesisedExpressionSyntax{
			OpenParenthesisToken:  left,
			Expression:            expression,
			CloseParenthesisToken: right,
		}
	}
	numberToken := p.match(NumberToken)
	return NumberExpressionSyntax{Token: numberToken}
}

func (p *Parser) parseTerm() ExpressionSyntax {
	left := p.parseFactor()
	for p.Current().Kind().isBinaryTermOperator() {
		operatorToken := p.NextToken()
		right := p.parseFactor()
		left = BinaryExpressionSyntax{
			Left:     left,
			Operator: operatorToken,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseFactor() ExpressionSyntax {
	left := p.parsePrimaryExpression()
	for p.Current().Kind().isBinaryFactorOperator() {
		operatorToken := p.NextToken()
		right := p.parsePrimaryExpression()
		left = BinaryExpressionSyntax{
			Left:     left,
			Operator: operatorToken,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) Parse() SyntaxTree {
	expression := p.parseTerm()
	endOfFileToken := p.match(EndOfFileToken)
	return SyntaxTree{
		diagnostics:    p.diagnostics,
		Root:           expression,
		endOfFileToken: endOfFileToken,
	}
}

func (p Parser) Peek(offset int) SyntaxToken {
	index := p.position + offset
	if index >= len(p.Tokens) {
		return p.Tokens[len(p.Tokens)-1]
	}
	return p.Tokens[index]
}

func (p Parser) Current() SyntaxToken {
	return p.Peek(0)
}
