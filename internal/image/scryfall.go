package image

import (
	"context"
	"fmt"
	"strings"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/gautierenaud/gocards/internal/log"
)

type Scryfall struct {
	log    log.Logger
	client scryfall.Client
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

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	resp, err := s.client.SearchCards(ctx, toQuery(parameters), sco)
	if err != nil {
		return "", err
	}

	if len(resp.Cards) == 0 {
		return "", nil
	}

	// TODO keep the alternatives in memory so the user can select the right one?
	if resp.Cards[0].ImageURIs == nil {
		return "", nil
	}

	return resp.Cards[0].ImageURIs.Normal, nil
}

func toQuery(p *Params) string {
	query := ""

	for k, v := range p.Parameters {
		switch k {
		case nameField:
			query += fmt.Sprintf("%s ", v)
		case setField:
			query += fmt.Sprintf("s:%s ", v)
		case setNumberField:
			query += fmt.Sprintf("cn:%s ", v)
		default:
			panic("unsupported parameter:" + k)
		}
	}

	return strings.Trim(query, " ")
}
