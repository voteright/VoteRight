package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/voteright/voteright/config"
)

var rootCmd = &cobra.Command{
	Use:   "voteright",
	Short: "Run electronic elections the right way.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VoteRight server starting")

		_, err := config.New()

		if err != nil {
			fmt.Println(err)
		}

	},
}

// Execute makes the root command runnable.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
