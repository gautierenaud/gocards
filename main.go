package main

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/gautierenaud/gocards/cmd"
	"github.com/gautierenaud/gocards/internal/config"
	"github.com/gautierenaud/gocards/internal/oracle"
	"github.com/gautierenaud/gocards/internal/store"
)

//go:embed all:frontend/dist
var assets embed.FS

var rootCmd = &cobra.Command{
	Use:          "gocard",
	Short:        "cli interface for gocard",
	SilenceUsage: true,
	RunE:         exec,
}

func main() {
	rootCmd.AddCommand(cmd.ImportCommand())

	rootCmd.Execute()
}

func exec(cmd *cobra.Command, args []string) error {
	// Create dependencies
	conf, err := config.LoadFromFile()
	if err != nil {
		return err
	}

	s, err := store.NewSQLiteStore(conf.Env.ConfigFolder)
	if err != nil {
		return err
	}

	fetcher, err := oracle.NewScryfall()
	if err != nil {
		return err
	}

	// Create an instance of the app structure
	app := NewApp(conf, s, fetcher)

	// Create application with options
	return wails.Run(&options.App{
		Title:  "gocards",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []any{
			app,
		},
	})
}
