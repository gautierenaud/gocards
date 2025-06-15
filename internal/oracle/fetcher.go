package oracle

import (
	"context"

	"github.com/gautierenaud/gocards/internal/models"
)

type Fetcher interface {
	GetImage(ctx context.Context, params ...Param) (string, error)
	GetSets(ctx context.Context) ([]models.Set, error)
	GetCards(ctx context.Context, params ...Param) ([]models.Card, error) // TODO think about pagination
}

type Params struct {
	Parameters map[string]any
}

type Param func(*Params)

const (
	nameField      = "name"
	setField       = "set"
	setNumberField = "set_number"
	languageField  = "language"
)

func WithName(name string) Param {
	return func(p *Params) {
		if name != "" {
			p.Parameters[nameField] = name
		}
	}
}

func WithSet(set string) Param {
	return func(p *Params) {
		if setField != "" {
			p.Parameters[setField] = set
		}
	}
}

func WithSetNumber(setNumber string) Param {
	return func(p *Params) {
		if setNumber != "" {
			p.Parameters[setNumberField] = setNumber
		}
	}
}

func WithLanguage(language string) Param {
	return func(p *Params) {
		if language != "" {
			p.Parameters[languageField] = language
		}
	}
}
