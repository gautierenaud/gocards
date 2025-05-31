package store_test

import "github.com/gautierenaud/gocards/internal/store"

// Check that SQLite implements Store interface
var _ store.Store = &store.SQLite{}
