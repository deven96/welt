package diagnostic

import (
	"fmt"
	"reflect"
)

type TextSpan struct {
	Start  int
	Length int
}

type Diagnostic struct {
	Span    TextSpan
	Message string
}

func (d Diagnostic) String() string {
	return d.Message
}

type DiagnosticsBag []Diagnostic

func (d *DiagnosticsBag) Report(span TextSpan, message string) {
	single := Diagnostic{span, message}
	*d = append(*d, single)
}

func (d *DiagnosticsBag) ReportInvalidNumber(span TextSpan, number string, typ reflect.Type) {
	message := fmt.Sprintf("The number %s isn't valid type %s.", number, typ)
	d.Report(span, message)
}

func (d *DiagnosticsBag) ReportBadCharacter(position int, character byte) {
	message := fmt.Sprintf("Bad character input: %b.", character)
	d.Report(TextSpan{position, 1}, message)
}

func (d *DiagnosticsBag) ReportUnexpectedToken(span TextSpan, expected, retrieved string) {
	message := fmt.Sprintf("Unexpected token: expected <%s> got <%s>.", expected, retrieved)
	d.Report(span, message)
}

func (d *DiagnosticsBag) ReportUndefinedUnaryOperator(span TextSpan, operator string, operandType reflect.Type) {
	message := fmt.Sprintf("Unary operator %s is not defined for %s.", operator, operandType)
	d.Report(span, message)
}

func (d *DiagnosticsBag) ReportUndefinedBinaryOperator(span TextSpan, operator string, left, right reflect.Type) {
	message := fmt.Sprintf("Binary operator %s is not defined for (%s, %s).", operator, left, right)
	d.Report(span, message)
}

func (d *DiagnosticsBag) AddBag(newBag DiagnosticsBag) {
	*d = append(*d, newBag...)
}

func (d DiagnosticsBag) ToString() []string {
	errors := []string{}
	for _, e := range d {
		errors = append(errors, fmt.Sprintf("[ERROR] %s", e))
	}
	return errors
}
