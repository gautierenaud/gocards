package oracle

import (
	"context"
	"fmt"
	"strings"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/pkg/errors"
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

func (s *Scryfall) GetSets(ctx context.Context) ([]models.Set, error) {
	sets, err := s.client.ListSets(ctx)
	if err != nil {
		return nil, err
	}

	codes := make([]models.Set, 0, len(sets))
	for _, set := range sets {
		codes = append(codes, models.Set{
			Name: set.Name,
			Code: set.Code,
		})
	}

	return codes, nil
}

func (s *Scryfall) GetCards(ctx context.Context, params ...Param) ([]models.Card, error) {
	// TODO convert parameters to scryfall compatible ones

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}

	cards, err := s.client.SearchCards(ctx, "", sco) // TODO replace query with one generated from params
	if err != nil {
		return nil, errors.Wrapf(err, "could not retrieve cards for set %s", "")
	}

	res := make([]models.Card, 0, len(cards.Cards))

	// TODO convert cards to internal format

	return res, nil
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
