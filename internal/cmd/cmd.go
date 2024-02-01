package cmd

import (
	"github.com/felipereyel/thor/internal/config"

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

	return startFiber(cfg)
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
