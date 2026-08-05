package store

import "time"

// Store is the key-value engine. YOU IMPLEMENT per specs/v1_store.md.
type Store struct {
	// fields yours
}

func New(maxMemoryBytes int) *Store {
	return &Store{}
}

func (s *Store) Set(key, value string, ttl time.Duration) error {
	return errUnimplemented
}

func (s *Store) Get(key string) (string, bool) {
	return "", false
}

func (s *Store) Del(key string) bool {
	return false
}

var errUnimplemented = errString("store not implemented")

type errString string

func (e errString) Error() string { return string(e) }
