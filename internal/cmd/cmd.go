package cmd

import (
	"fmt"

	"github.com/felipereyel/thor/internal/config"
	"github.com/felipereyel/thor/internal/services"

	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetServerConfigs()
	if err != nil {
		return err
	}

	if err := migrateDb(cfg); err != nil {
		return err
	}

	svcs, err := services.NewServices(cfg)
	if err != nil {
		return err
	}

	if err := recoverTorrents(svcs); err != nil {
		fmt.Printf("error recovering torrents: %s\n", err)
	}

	return startFiber(svcs)
}

var rootCmd = &cobra.Command{
	Use:   "thor",
	Short: "start thor server",
	RunE:  Start,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
