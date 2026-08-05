package resp

import (
	"fmt"
	"io"
)

type Kind int

const (
	SimpleString Kind = iota
	Error
	Integer
	BulkString
	Array
)

type Value struct {
	Kind   Kind
	Str    string
	Int    int64
	Bulk   []byte // nil => null bulk
	Array  []Value
	IsNull bool
}

// Decode reads one RESP value. YOU IMPLEMENT.
func Decode(r io.Reader) (Value, error) {
	return Value{}, fmt.Errorf("resp.Decode not implemented")
}

// Encode writes one RESP value. YOU IMPLEMENT.
func Encode(w io.Writer, v Value) error {
	return fmt.Errorf("resp.Encode not implemented")
}
