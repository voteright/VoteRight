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

	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initDatabaseCmd represents the initDatabase command
var initDatabaseCmd = &cobra.Command{
	Use:   "initDatabase",
	Short: "Initialize the database for use with the voteright server",

	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}
		err := viper.Unmarshal(&cfg)

		if err != nil {
			fmt.Println("Failed to unmarshal configuration!")
			return
		}
		d, err := database.New(&cfg)
		if err != nil {
			return
		}
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS voter (id INTEGER PRIMARY KEY, name VARCHAR(255), cohort INTEGER)")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS votes (hash INTEGER PRIMARY KEY, candidateid INTEGER)")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS cohorts (id INTEGER PRIMARY KEY, name VARCHAR(255))")
		if err != nil {
			fmt.Println(err.Error())
		}
		// For the purposes of this class, for now, these values will be hardcoded, eventually we should allow configuration
		_, err = d.ExecStatement("INSERT INTO cohorts VALUES (0,'2019'),(1,'2020'), (2,'2021'), (3,'2022')")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS candidates (id INTEGER PRIMARY KEY, candidateid INTEGER, cohort INTEGER, FOREIGN KEY (cohort) REFERENCES cohorts(id))")
		if err != nil {
			fmt.Println(err.Error())
		}
		d.ExecStatement("INSERT INTO voter values (1, \"TEST\")")

	},
}

func init() {
	rootCmd.AddCommand(initDatabaseCmd)
}
