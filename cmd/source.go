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
	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Get a link to the source code",
	Long: `Just a quick and easy way to get a link
to the source code of QSR.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(NewMessage(chalk.Blue, "Go check out the source @").
			ThenColorStyle(chalk.Blue, chalk.Underline, "https://github.com/GreatGodApollo/qsr").Build())
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
