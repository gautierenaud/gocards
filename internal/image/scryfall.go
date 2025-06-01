package image

import (
	"context"
	"fmt"
	"strings"

	scryfall "github.com/BlueMonday/go-scryfall"
)

type Scryfall struct {
	client scryfall.Client
}

func NewScryfall() (*Scryfall, error) {
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, err
	}

	return &Scryfall{
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
		if strings.HasPrefix(err.Error(), "not_found") {
			// fallback to english
			WithLanguage("en")(parameters)
			resp, err = s.client.SearchCards(ctx, toQuery(parameters), sco)
		}

		// if the error is still present return it
		if err != nil {
			return "", err
		}
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
		case languageField:
			query += fmt.Sprintf("lang:%s ", v)
		default:
			// do nothing, we might have option that could be used in other fetchers
		}
	}

	return strings.Trim(query, " ")
}
