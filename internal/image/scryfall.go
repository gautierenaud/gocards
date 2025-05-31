package image

import (
	"context"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/gautierenaud/gocards/internal/log"
	"golang.org/x/time/rate"
)

type Scryfall struct {
	log     log.Logger
	client  scryfall.Client
	limiter *rate.Limiter
}

func NewScryfall(log log.Logger) (*Scryfall, error) {
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, err
	}

	return &Scryfall{
		log:    log,
		client: *client,
	}, nil
}

func (s *Scryfall) GetImage(ctx context.Context, params ...Param) (string, error) {
	parameters := &Params{
		Parameters: make(map[string]any),
	}
	for _, param := range params {
		param(parameters)
	}

	s.log.Infof("Retrieving image for %s", parameters.Parameters["cardName"])

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	resp, err := s.client.SearchCards(ctx, parameters.Parameters["cardName"].(string), sco)
	if err != nil {
		return "", err
	}

	if len(resp.Cards) == 0 {
		return "", nil
	}

	if resp.Cards[0].ImageURIs == nil {
		return "", nil
	}

	return resp.Cards[0].ImageURIs.Normal, nil
}
