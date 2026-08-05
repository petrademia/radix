package store_test

import (
	"testing"
	"time"

	"github.com/petrademia/radix/internal/store"
)

func TestSetGet(t *testing.T) {
	s := store.New(1 << 20)
	if err := s.Set("a", "b", 0); err != nil {
		t.Fatal(err)
	}
	v, ok := s.Get("a")
	if !ok || v != "b" {
		t.Fatalf("got %q %v", v, ok)
	}
}

func TestTTLExpires(t *testing.T) {
	s := store.New(1 << 20)
	_ = s.Set("a", "b", 10*time.Millisecond)
	time.Sleep(20 * time.Millisecond)
	if _, ok := s.Get("a"); ok {
		t.Fatal("expected expiry")
	}
}
