package syntax

import (
	"github.com/deven96/welt/diagnostic"
)

type Parser struct {
	Tokens      []SyntaxToken
	position    int
	diagnostics diagnostic.DiagnosticsBag
}

func NewParser(text string) Parser {
	p := Parser{}
	lexer := lexer{Text: text}
	tokens := []SyntaxToken{}
	for {
		token := lexer.Lex()
		if token.Kind_ != WhitespaceToken && token.Kind_ != BadToken {
			tokens = append(tokens, token)
		}
		if token.Kind_ == EndOfFileToken {
			break
		}
	}
	p.Tokens = tokens
	p.diagnostics.AddBag(lexer.diagnostics)
	return p
}

func (p *Parser) NextToken() SyntaxToken {
	current := p.Current()
	p.position++
	return current
}

func (p *Parser) matchToken(kind SyntaxKind) SyntaxToken {
	if p.Current().Kind_ == kind {
		return p.NextToken()
	}
	p.diagnostics.ReportUnexpectedToken(p.Current().Span(), kind.String(), p.Current().Kind().String())
	return SyntaxToken{
		Kind_:    kind,
		position: p.Current().position,
	}
}

func (p *Parser) Parse() SyntaxTree {
	expression := p.parseAssignmentExpression()
	endOfFileToken := p.matchToken(EndOfFileToken)
	return SyntaxTree{
		diagnostics:    p.diagnostics,
		Root:           expression,
		endOfFileToken: endOfFileToken,
	}
}

func (p *Parser) parsePrimaryExpression() ExpressionSyntax {
	currentKind := p.Current().Kind()
	switch currentKind {
	case OpenParenthesisToken:
		left := p.NextToken()
		expression := p.parseAssignmentExpression()
		right := p.matchToken(CloseParenthesisToken)
		return ParenthesisedExpressionSyntax{
			OpenParenthesisToken:  left,
			Expression:            expression,
			CloseParenthesisToken: right,
		}
	case TrueKeyWord, FalseKeyWord:
		keyWordToken := p.NextToken()
		value := currentKind == TrueKeyWord
		return LiteralExpressionSyntax{
			LiteralToken: keyWordToken,
			Value:        value,
		}
	case QuotedIdentifierToken:
		token := p.NextToken()
		return QuotedExpressionSyntax{QuotedIdentifier: token}
	case IdentifierToken:
		identifier := p.NextToken()
		return NameExpressionSyntax{Identifier: identifier}
	default:
		numberToken := p.matchToken(NumberToken)
		if numberToken.Value != nil {
			value := numberToken.Value.(int)
			return LiteralExpressionSyntax{LiteralToken: numberToken, Value: value}
		} else {
			return LiteralExpressionSyntax{LiteralToken: numberToken}
		}
	}
}

func (p *Parser) parseAssignmentExpression() ExpressionSyntax {
	if p.Current().Kind() == IdentifierToken && p.LookAhead().Kind() == EqualsToken {
		identifier := p.matchToken(IdentifierToken)
		operator := p.matchToken(EqualsToken)
		right := p.parseAssignmentExpression()
		return AssignmentExpressionSyntax{
			Identifier:  identifier,
			EqualsToken: operator,
			Expression:  right,
		}
	}
	return p.parseBinaryExpression(0)
}

func (p *Parser) parseBinaryExpression(parentPrecendence int) ExpressionSyntax {
	var left ExpressionSyntax
	unaryOperatorPrecedence := p.Current().Kind().getUnaryOperatorPrecedence()
	if unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecendence {
		operator := p.NextToken()
		operand := p.parseBinaryExpression(unaryOperatorPrecedence)
		left = UnaryExpressionSyntax{
			Operator: operator,
			Operand:  operand,
		}
	} else {
		left = p.parsePrimaryExpression()
	}
	for {
		precedence := p.Current().Kind().getBinaryOperatorPrecedence()
		if precedence == 0 || precedence <= parentPrecendence {
			break
		}
		operatorToken := p.NextToken()
		right := p.parseBinaryExpression(precedence)
		left = BinaryExpressionSyntax{
			Left:     left,
			Operator: operatorToken,
			Right:    right,
		}
	}
	return left
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
func (p Parser) LookAhead() SyntaxToken {
	return p.Peek(1)
}
