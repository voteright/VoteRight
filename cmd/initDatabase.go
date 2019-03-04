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
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS voters (id INTEGER PRIMARY KEY, name VARCHAR(255), cohort INTEGER)")
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
		_, err = d.ExecStatement("CREATE TABLE IF NOT EXISTS candidates (id INTEGER PRIMARY KEY, name VARCHAR(255), cohort INTEGER, FOREIGN KEY (cohort) REFERENCES cohorts(id))")
		if err != nil {
			fmt.Println(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(initDatabaseCmd)
}
