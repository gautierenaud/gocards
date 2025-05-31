package importer

import (
	"bufio"
	"os"
	"regexp"

	"github.com/gautierenaud/gocards/internal/log"
	"github.com/gautierenaud/gocards/internal/models"
	"github.com/pkg/errors"
)

func Import(log log.Logger, path string) ([]*models.Card, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not open file")
	}
	defer file.Close()

	cards := make([]*models.Card, 0, 100)

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`(\d+)\s(.*)\s\(([0-9A-Z]{3,4})\)\s(\d+s?)\s?(F?)`) // 1 Chord of Calling (SLD) 1595 F
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) == 0 {
			return nil, errors.New("unsupported format: " + line)
		}

		// TODO store other information as well
		cards = append(cards, &models.Card{
			Name:      matches[2],
			Set:       matches[3],
			SetNumber: matches[4],
		})
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrap(err, "could not read file")

	}

	return cards, nil
}
