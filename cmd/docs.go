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
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

var dir string

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs [directory]",
	Short: "Documentation Generator",
	Long: `A super quick command to generate documentation for Quick Script Runner`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.Mkdir(dir, os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}
		}
		err := doc.GenMarkdownTree(rootCmd, dir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringVarP(&dir, "directory", "d", "./", "Help message for toggle")
}
