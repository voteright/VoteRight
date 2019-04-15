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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/voteright/voteright/api"
	"github.com/voteright/voteright/config"
	"github.com/voteright/voteright/database"
	"github.com/voteright/voteright/election"
)

var cfgFile string
var portOverride string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "voteright",
	Short: "A verifiable voting suite, primary server",
	Long: `This is the primary voting server that serves the voting booth, and sends votes to a verification cluster.
	
	This should be run by a trusted party, and configured in config.json as such:
	{Fill this in}`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}
		err := viper.Unmarshal(&cfg)
		fmt.Println("config", cfg.VerificationServers)
		// d, err := database.New(&cfg)
		// if err != nil {
		// 	return
		// }

		if portOverride != "" {
			cfg.ListenURL = "0.0.0.0:" + portOverride
		}
		d := database.StormDB{
			File: cfg.DatabaseFile,
		}
		d.Connect()

		if err != nil {
			fmt.Println("Failed to unmarshal configuration!")
			return
		}

		e := election.New(&d, false, cfg.VerificationServers)

		api := api.New(&cfg, e, &d)
		api.Serve()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./voteright.json)")
	rootCmd.Flags().StringVarP(&portOverride, "port", "p", "8080", "override server port")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		viper.AddConfigPath(dir)
		viper.SetConfigType("json")
		viper.SetConfigName("voteright")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Print(err)
	}

}
