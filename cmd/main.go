package main

import (
	"github.com/axernel-labs/stackdome/cmd/migratecmd"
	"github.com/axernel-labs/stackdome/cmd/servecmd"
	"github.com/axernel-labs/stackdome/pkg/logger"
	"github.com/spf13/cobra"
)

var log = logger.NewLogger()

func main() {
	rootCmd := &cobra.Command{
		Use:  "stackdome-server",
		Long: "stackdome-server",
	}

	serveCmd := servecmd.NewServeCommand()
	migrateCmd := migratecmd.NewMigrateCommand()
	rootCmd.AddCommand(serveCmd, migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error running command: %v", err)
	}
}
