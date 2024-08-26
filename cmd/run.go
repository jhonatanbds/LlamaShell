/*
Copyright © 2024 jhonatan <jhonatanbds97@gmail.com>
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/jhonatanbds/LlamaShell/internal/run"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run.Run(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
