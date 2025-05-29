package store_test

import "gocards/backend/internal/store"

// Check that SQLite implements Store interface
var _ store.Store = &store.SQLite{}
