package main

import (
	"github.com/spf13/cobra"
)

// funcsCmd represents the funcs command
var groupCmd = &cobra.Command{
	Use: "group",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(funcsCmd)
}
