package cmd

import (
	"github.com/keigo-saito0602/go_learning_for_poke_api/cmd/migrate"
	"github.com/keigo-saito0602/go_learning_for_poke_api/cmd/server"
	"github.com/spf13/cobra"
)

func Run() int {
	if err := rootCommand().Execute(); err != nil {
		return 1
	}
	return 0
}

func rootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "go_learning_for_poke_api",
		Short: "go learning for poke api CLI",
	}

	for _, g := range commandGroups() {
		root.AddGroup(g)
	}

	root.AddCommand(
		server.ServerCommand(),
		migrate.MigrateCommand(),
		newServeCmd(),
	)

	return root
}
