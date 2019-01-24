package errors

import "fmt"

type InputError struct {
	input string
	err   error
}

func NewInputError(input string, err error) error {
	return &InputError{input, err}
}

func (i *InputError) Error() string {
	return fmt.Sprintf("Invalid input %s, error: %s", i.input, i.err)
}

type ParseError struct {
	cursor    int
	character byte
}

func NewParseError(character byte, cursor int) *ParseError {
	return &ParseError{cursor, character}
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("invalid character '%s' looking for at pos %d of value", p.character, p.cursor)
}

type NumberOverflowError struct {
	index int
	typ   string
}

func NewInt8OverflowError(index int) error {
	return &NumberOverflowError{index, "Int8"}
}

func NewInt16OverflowError(index int) error {
	return &NumberOverflowError{index, "Int16"}
}

func NewInt32OverflowError(index int) error {
	return &NumberOverflowError{index, "Int32"}
}

func NewInt64OverflowError(index int) error {
	return &NumberOverflowError{index, "Int64"}
}

func NewUint8OverflowError(index int) error {
	return &NumberOverflowError{index, "Uint8"}
}

func NewUint16OverflowError(index int) error {
	return &NumberOverflowError{index, "Uint16"}
}

func NewUint32OverflowError(index int) error {
	return &NumberOverflowError{index, "Uint32"}
}

func NewUint64OverflowError(index int) error {
	return &NumberOverflowError{index, "Uint64"}
}

func NewFloat32OverflowError(index int) error {
	return &NumberOverflowError{index, "Float32"}
}

func NewFloat64OverflowError(index int) error {
	return &NumberOverflowError{index, "Float64"}
}

func (n *NumberOverflowError) Error() string {
	return fmt.Sprintf("invalid json, %s overflow at pos %d", n.typ, n.index)
}
