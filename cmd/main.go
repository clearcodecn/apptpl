package main

import (
	"{{ .pkgName}}/cmd/server"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:        "cli",
		Aliases:    nil,
		SuggestFor: nil,
	}
)

func init() {
	rootCmd.AddCommand(
		server.ServerCmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
