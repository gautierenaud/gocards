package oracle_test

import "github.com/gautierenaud/gocards/internal/oracle"

// Check that we implement the Fetcher interface.
var _ oracle.Fetcher = &oracle.Scryfall{}
