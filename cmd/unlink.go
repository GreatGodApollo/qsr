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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var unlinkCmd = &cobra.Command{
	Use:   "unlink <alias>",
	Short: "Remove a gist or file alias",
	Long: `Remove a gist or file alias from your config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if viper.AllSettings()[args[0]] != nil {
			err := Unset(args[0])
			if CheckError(err) {
				return
			}
			fmt.Println(NewMessage(chalk.Green, "Successfully deleted gist alias").
				ThenColor(chalk.Cyan, args[0]).Build())
		} else {
			fmt.Println(NewMessage(chalk.Red, "Gist alias").
				ThenColor(chalk.Cyan, args[0]).
				ThenColor(chalk.Red, "does not exist!").Build())
		}
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}

func Unset(key string) error {
	configMap := viper.AllSettings()
	delete(configMap, key)
	encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
	err := viper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}
	return viper.WriteConfig()
}