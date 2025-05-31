package image_test

import "github.com/gautierenaud/gocards/internal/image"

// Check that we implement the Fetcher interface.
var _ image.Fetcher = &image.Scryfall{}
