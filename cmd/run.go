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
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/ttacon/chalk"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var justPrint bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [gist id] [filename]",
	Short: "Run a remote gist",
	Long: `Quick Script Runner is a command line utility that allows you to run gists
with a single command.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		gcli := github.NewClient(&http.Client{})
		gist, _, err := gcli.Gists.Get(context.Background(), args[0])
		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That gist couldn't be found!"))
			} else if strings.Contains(err.Error(), "no such host") {
				fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("It looks like you don't have an internet connection!"))
			} else {
				fmt.Println(err.Error())
			}
			return
		}
		file := gist.Files[github.GistFilename(args[1])]
		if justPrint {
			fmt.Println(chalk.Cyan.Color("[QSR]"), "FileType:",chalk.Blue.Color(file.GetLanguage()))
			fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Magenta.Color(file.GetContent()))
		} else {

			if file.GetSize() > 0 {
				startingDirectory, err := os.Getwd()
				if CheckError(err) {
					return
				}

				if runtime.GOOS == "windows" {
					err = os.Chdir(os.Getenv("TEMP"))
				} else if runtime.GOOS == "linux" {
					if _, err := os.Stat("/tmp/qsr/"); os.IsNotExist(err) {
						os.Mkdir("/tmp/qsr/", os.ModeDir)
					}
				}
				if CheckError(err) {
					return
				}
				switch file.GetLanguage() {
				case "Batchfile":
					{
						if runtime.GOOS == "windows" {
							err = DownloadFile("tmp.bat", file.GetRawURL())
							if CheckError(err) {
								return
							}
							RunCommand("tmp.bat", "", "")
							err = os.Remove("tmp.bat")
							if CheckError(err) {
								return
							}
							break
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't " +
								"supported on your system!"))
						}
					}
				case "Shell":
					{
						if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
							err = DownloadFile("tmp.sh", file.GetRawURL())
							if CheckError(err) {
								return
							}
							RunCommand("chmod", "+x", "tmp.sh")
							RunCommand("./tmp.sh", "", "")
							err = os.Remove("tmp.sh")
							if CheckError(err) {
								return
							}
							break
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't " +
								"supported on your system!"))
						}
					}
				case "Go":
					{
						err = DownloadFile("tmp.go", file.GetRawURL())
						if CheckError(err) {
							return
						}
						RunCommand("go", "run", "tmp.go")
						err = os.Remove("tmp.go")
						if CheckError(err) {}
						break
					}
				case "JavaScript":
					{
						err = DownloadFile("tmp.js", file.GetRawURL())
						if CheckError(err) {
							return
						}
						RunCommand("node", "tmp.js", "")
						err = os.Remove("tmp.js")
						if CheckError(err) {
							return
						}
						break
					}
				case "Ruby":
					{
						err = DownloadFile("tmp.rb", file.GetRawURL())
						if CheckError(err) {
							return
						}
						RunCommand("ruby", "tmp.rb", "")
						err = os.Remove("tmp.rb")
						if CheckError(err) {
							return
						}
						break
					}
				default:
					{
						fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't supported!"))
						break
					}
				}
				err = os.Chdir(startingDirectory)
				if CheckError(err) {
					return
				}
			} else {
				fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That files doesn't exist!"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&justPrint, "print", "p", false, "Print out the script instead of running it")
}
