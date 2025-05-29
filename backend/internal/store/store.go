package store

import "gocards/backend/internal/data"

type Store interface {
	Setup(setupFolder string) error
	All() ([]data.Card, error)
}
