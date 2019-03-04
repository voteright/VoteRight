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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"
)

// importdbCmd represents the importdb command
var importdbCmd = &cobra.Command{
	Use: "importdb [filename]",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments for import (Requires filename)")
			return
		}

		fmt.Printf("Importing database from %s\n", args[0])

		cfg := config.Config{}
		err := viper.Unmarshal(&cfg)

		if err != nil {
			fmt.Println("Failed to unmarshal configuration!")
			return
		}
		d, err := database.New(&cfg)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_ = d
		dump := database.Dump{}
		bytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("Failed to read file", err.Error())
			return
		}
		err = json.Unmarshal(bytes, &dump)
		if err != nil {
			if err != nil {
				fmt.Println("Invalid format", err.Error())
				return
			}
		}
		fmt.Println(dump)

		err = d.StoreCohorts(dump.Cohorts)
		if err != nil {
			fmt.Println("Error importing cohorts:", err.Error())
			fmt.Println("Was the database empty? Continuing")
		}

		err = d.StoreCandidates(dump.Candidates)
		if err != nil {
			fmt.Println("Error importing candidates:", err.Error())
			fmt.Println("Was the database empty? Continuing")

		}

		err = d.StoreVoters(dump.Voters)
		if err != nil {
			fmt.Println("Error importing voters:", err.Error())
			fmt.Println("Was the database empty? Continuing")

		}

		for _, vote := range dump.Votes {
			err = d.StoreVote(vote)
			if err != nil {
				fmt.Println("Error importing votes:", err.Error())
				fmt.Println("Was the database empty? Continuing")

			}

		}

		fmt.Println("Done")

	},
}

func init() {
	rootCmd.AddCommand(importdbCmd)
}
