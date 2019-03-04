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

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// exportdbCmd represents the exportdb command
var exportdbCmd = &cobra.Command{
	Use: "exportdb",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments for export (Requires filename)")
			return
		}

		fmt.Printf("Exporting database to %s\n", args[0])

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

		voters, err := d.GetAllCohorts()
		if err != nil {
			fmt.Println(err.Error())

			return
		}
		_ = voters
		dump := database.Dump{}
		dump.Voters, err = d.GetAllVoters()
		if err != nil {
			fmt.Println("Error Exporting Voters")
			return
		}
		dump.Candidates, err = d.GetAllCandidates()
		if err != nil {
			fmt.Println("Error Exporting Candidates")
			return
		}
		dump.Cohorts, err = d.GetAllCohorts()
		if err != nil {
			fmt.Println("Error Exporting Cohorts")
			return
		}
		dump.Votes, err = d.GetAllVotes()
		if err != nil {
			fmt.Println("Error Exporting Votes")
			return
		}

		json, _ := json.Marshal(dump)
		err = ioutil.WriteFile(args[0], []byte(json), 0644)
		if err != nil {
			fmt.Println("Error writing to file")
		} else {
			fmt.Println("Done")
		}
	},
}

func init() {
	rootCmd.AddCommand(exportdbCmd)

}
