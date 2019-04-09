// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/election"
	"github.com/voteright/voteright/verificationapi"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}
		err := viper.Unmarshal(&cfg)

		d, err := database.New(&cfg)
		if err != nil {
			return
		}
		if err != nil {
			fmt.Println("Failed to unmarshal configuration!")
			return
		}

		e := election.New(d)

		api := verificationapi.New(&cfg, e, d)
		api.Serve()
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

}
