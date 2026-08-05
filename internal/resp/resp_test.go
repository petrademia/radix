package resp_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/petrademia/radix/internal/resp"
)

func TestDecodeSimpleString(t *testing.T) {
	v, err := resp.Decode(strings.NewReader("+OK\r\n"))
	if err != nil {
		t.Fatal(err)
	}
	if v.Kind != resp.SimpleString || v.Str != "OK" {
		t.Fatalf("got %+v", v)
	}
}

func TestRoundTripBulk(t *testing.T) {
	var buf bytes.Buffer
	in := resp.Value{Kind: resp.BulkString, Bulk: []byte("hello")}
	if err := resp.Encode(&buf, in); err != nil {
		t.Fatal(err)
	}
	out, err := resp.Decode(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if string(out.Bulk) != "hello" {
		t.Fatalf("got %+v", out)
	}
}
