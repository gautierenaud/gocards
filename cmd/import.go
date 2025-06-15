package cmd

import (
	"log/slog"

	"github.com/mdouchement/logger"
	"github.com/spf13/cobra"

	"github.com/gautierenaud/gocards/internal/config"
	"github.com/gautierenaud/gocards/internal/importer"
	"github.com/gautierenaud/gocards/internal/oracle"
	"github.com/gautierenaud/gocards/internal/store"
)

func ImportCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "import <filename>",
		Short: "import cardlist from a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.WrapSlog(slog.Default())
			ctx := logger.WithLogger(cmd.Context(), log)

			conf, err := config.LoadFromFile()
			if err != nil {
				return err
			}

			file := args[0]
			cards, err := importer.Import(ctx, file)
			if err != nil {
				return err
			}

			log.Infof("We have %d images to retrieve", len(cards))

			fetcher, err := oracle.NewScryfall(log)
			if err != nil {
				return err
			}

			// TODO retrieve image only if we don't have it in the database
			for _, card := range cards {
				image, err := fetcher.GetImage(ctx,
					oracle.WithName(card.Name),
					oracle.WithSet(card.Set),
					oracle.WithSetNumber(card.SetNumber),
					oracle.WithLanguage(conf.App.Language),
				)
				if err != nil {
					log.Errorf("Could not retrieve image for %s: %s", card.Name, err)
				}

				card.ImagePath = image
			}

			store, err := store.NewSQLiteStore(conf.Env.ConfigFolder)
			if err != nil {
				return err
			}

			err = store.Store(cards)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
