package image

import "context"

type Fetcher interface {
	GetImage(ctx context.Context, params ...Param) (string, error)
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
