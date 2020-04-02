/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// createCmd represents the create command
var linkCmd = &cobra.Command{
	Use:   "link <alias> <gist> [<file>]",
	Short: "Link a gist or file alias",
	Long: `Link a gist or file alias in your config.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			viper.Set(args[0], map[string]string{"gist": args[1], "file": ""})
			err := viper.WriteConfig()
			if CheckError(err) {return}
			fmt.Println(NewMessage(chalk.Green, "Successfully wrote gist alias").
				ThenColor(chalk.Cyan, args[0]).
				ThenColor(chalk.Green, "as").
				ThenColor(chalk.Cyan, args[1]))
		} else {
			viper.Set(args[0], map[string]string{"gist": args[1], "file": args[2]})
			err := viper.WriteConfig()
			if CheckError(err) {return}
			fmt.Println(NewMessage(chalk.Green, "Successfully wrote gist alias").
				ThenColor(chalk.Cyan, args[0]).
				ThenColor(chalk.Green, "as").
				ThenColor(chalk.Cyan, args[1]).
				ThenColor(chalk.Green,"with file").
				ThenColor(chalk.Cyan, args[2]))
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
