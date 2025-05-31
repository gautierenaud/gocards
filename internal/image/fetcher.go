package image

import "context"

type Params struct {
	Parameters map[string]any
}

type Param func(*Params)

type Fetcher interface {
	GetImage(ctx context.Context, params ...Param) (string, error)
}
