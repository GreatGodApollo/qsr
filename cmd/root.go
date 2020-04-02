/*
 *     Quick Script Runner: A quick and easy way to run gists
 *     Copyright © 2020 Brett Bender
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
	"os"
	"strings"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qsr <command> [<arguments>]",
	Short: "A quick and easy way to run gists",
	Long: `Quick Script Runner is a command line utility that allows you to run gists
with a single command.`,
	Version: "0.9.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "unknown command") {
			return
		}
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qsr.json)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if CheckError(err) {os.Exit(1)}

		viper.AddConfigPath(home)
		viper.SetConfigName(".qsr")
		viper.SetConfigType("json")
		if createFileIfNotExist(home + "/", ".qsr.json") != nil {os.Exit(1)}

	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(NewMessage(chalk.Blue, "Using config file:").ThenColor(chalk.Green, viper.ConfigFileUsed()))
	}
	SetGists()
	if CheckError(viper.WriteConfig()) {os.Exit(1)}
}

func createFileIfNotExist(dir, file string) error {
	if _, err := os.Stat(dir + file); err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(dir + file)
			if CheckError(err) {return err}
			err = f.Close()
			if CheckError(err) {return err}
		} else {
			return err
		}
	}
	return nil
}
