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
	"github.com/dixonwille/wmenu/v5"
	"github.com/google/go-github/github"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var justPrint bool
var yes bool
var dlpath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run (<gist_id> <file>|<alias> [<file>]) [<arguments>...]",
	Short: "Run a remote gist",
	Long: `Quick Script Runner is a command line utility that allows you to run gists
with a single command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, inargs []string) {
		var args = make([]string, 2)
		if viper.ConfigFileUsed() != "" {
			sub := viper.GetStringMapString(inargs[0])
			if len(sub) != 0 {
				if sub["gist"] != "" && sub["file"] != "" {
					args[0] = sub["gist"]
					args[1] = sub["file"]
					if len(inargs) > 1 {
						args = append(args, inargs[1:]...)
					}
				} else if sub["gist"] != "" && sub["file"] == "" {
					args[0] = sub["gist"]
					if len(inargs) > 1 {
						args[1] = inargs[1]
						args = append(args, inargs[2:]...)
					} else {
						fmt.Println(NewMessage(chalk.Red, "You must provide a file to run!"))
						return
					}
				} else {
					fmt.Println(NewMessage(chalk.Red, "Invalid alias configuration!"))
					return
				}
			} else {
				args = inargs
			}
		} else {
			args = inargs
		}
		var gcli *github.Client
		ctx := context.Background()
		if viper.GetString("token") != "" {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: viper.GetString("token")},
			)
			tc := oauth2.NewClient(ctx, ts)
			gcli = github.NewClient(tc)
			fmt.Println(NewMessage(chalk.Green, "Using authentication token"))
		} else {
			gcli = github.NewClient(&http.Client{})
		}

		gist, _, err := gcli.Gists.Get(ctx, args[0])
		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				fmt.Println(NewMessage(chalk.Red, "That gist couldn't be found!"))
			} else if strings.Contains(err.Error(), "no such host") {
				fmt.Println(NewMessage(chalk.Red, "It looks like you don't have an internet connection!"))
			} else if strings.Contains(err.Error(), "401 Bad credentials") {
				fmt.Println(NewMessage(chalk.Red, "Looks like your OAuth token is invalid or misconfigured!"))
			} else {
				fmt.Println(err.Error())
			}
			return
		}
		file := gist.Files[github.GistFilename(args[1])]
		if justPrint {
			fmt.Println(chalk.Cyan.Color("[QSR]"), "FileType:", chalk.Blue.Color(file.GetLanguage()))
			a := strings.Split(file.GetContent(), "\n")
			b := strings.Join(a, "\n"+chalk.Cyan.Color("[QSR] "))
			fmt.Println(chalk.Cyan.Color("[QSR]"), b)
		} else {

			if file.GetSize() > 0 {
				if !yes {
					menu := wmenu.NewMenu(NewMessage(chalk.Red, "Are you sure you want to run this script?").String())
					menu.IsYesNo(wmenu.DefN)
					menu.Action(verifyYes)
					err := menu.Run()
					if err != nil {
						fmt.Println(NewMessage(chalk.Red, err.Error()))
						return
					}
					if !yes {
						return
					}
				}

				home, err := homedir.Dir()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				directory := ""
				if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
					directory = home + "/.qsr/"
				} else if runtime.GOOS == "windows" {
					directory = home + "\\.qsr\\"
				}
				if _, err := os.Stat(directory); err != nil {
					if os.IsNotExist(err) {
						err = os.Mkdir(directory, os.ModeDir)
						if CheckError(err) {
							return
						}
					} else {
						if CheckError(err) {
							return
						}
					}
				}
				dlpath = os.Getenv(directory)

				switch file.GetLanguage() {
				case "Batchfile":
					{
						if runtime.GOOS == "windows" {
							err = DownloadFile(dlpath+"tmp.bat", file.GetRawURL())
							if CheckError(err) {
								return
							}
							err = RunCommand(dlpath+"tmp.bat", args[2:]...)
							CheckError(err)
							err = os.Remove(dlpath + "tmp.bat")
							if CheckError(err) {
								return
							}
							break
						} else {
							fmt.Println(NewMessage(chalk.Red, "That language isn't supported on your system!"))
							return
						}
					}
				case "Go":
					{
						err = DownloadFile(dlpath+"tmp.go", file.GetRawURL())
						if CheckError(err) {
							return
						}
						a := []string{"run", dlpath + "tmp.go"}
						a = append(a, args[2:]...)
						err = RunCommand("go", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(NewMessage(chalk.Red, "You can find instructions to install it here:").
								ThenColorStyle(chalk.Red, chalk.Underline, "https://golang.org"))
						}
						err = os.Remove(dlpath + "tmp.go")
						if CheckError(err) {
							return
						}
						break
					}
				case "JavaScript":
					{

						err = DownloadFile(dlpath+"tmp.js", file.GetRawURL())
						if CheckError(err) {
							return
						}
						a := []string{dlpath + "tmp.js"}
						a = append(a, args[2:]...)
						err = RunCommand("node", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(NewMessage(chalk.Red, "You can find instructions to install it here:").
								ThenColorStyle(chalk.Red, chalk.Underline, "https://nodejs.org"))
						}
						err = os.Remove(dlpath + "tmp.js")
						if CheckError(err) {
							return
						}
						break
					}
				case "Python":
					{
						err = DownloadFile(dlpath+"tmp.py", file.GetRawURL())
						if CheckError(err) {
							return
						}
						fst, err := getFirstLine(dlpath + "tmp.py")
						if CheckError(err) {
							return
						}
						if fst == "#!/usr/bin/python3" {
							a := []string{dlpath + "tmp.py"}
							a = append(a, args[2:]...)
							err = RunCommand("python3", a...)
							CheckError(err)
							if err != nil && strings.Contains(err.Error(), "executable file not found") {
								fmt.Println(NewMessage(chalk.Red, "You can find instructions to install it here:").
									ThenColorStyle(chalk.Red, chalk.Underline, "https://www.python.org"))
							}
						} else if fst == "#!/usr/bin/python" {
							a := []string{dlpath + "tmp.py"}
							a = append(a, args[2:]...)
							err = RunCommand("python", a...)
							CheckError(err)
							if err != nil && strings.Contains(err.Error(), "executable file not found") {
								fmt.Println(NewMessage(chalk.Red, "You can find instructions to install it here:").
									ThenColorStyle(chalk.Red, chalk.Underline, "https://www.python.org"))
							}
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("Cannot determine language!"))
						}
						err = os.Remove(dlpath + "tmp.py")
						if CheckError(err) {
							return
						}
						break
					}
				case "Ruby":
					{
						err = DownloadFile(dlpath+"tmp.rb", file.GetRawURL())
						if CheckError(err) {
							return
						}
						a := []string{dlpath + "tmp.rb"}
						a = append(a, args[2:]...)
						err = RunCommand("ruby", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(NewMessage(chalk.Red, "You can find instructions to install it here:").
								ThenColorStyle(chalk.Red, chalk.Underline, "https://ruby-lang.org"))
						}
						err = os.Remove(dlpath + "tmp.rb")
						if CheckError(err) {
							return
						}
						break
					}
				case "Shell":
					{
						if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
							err = DownloadFile(dlpath+"tmp.sh", file.GetRawURL())
							if CheckError(err) {
								return
							}
							err = RunCommand("chmod", "+x", dlpath+"tmp.sh")
							CheckError(err)
							err = RunCommand("./"+dlpath+"tmp.sh", args[2:]...)
							CheckError(err)
							err = os.Remove(dlpath + "tmp.sh")
							if CheckError(err) {
								return
							}
							break
						} else {
							fmt.Println(NewMessage(chalk.Red, "That language isn't supported on your system!"))
							return
						}
					}
				default:
					{
						fmt.Println(NewMessage(chalk.Red, "That language isn't supported!"))
						break
					}
				}
			} else {
				fmt.Println(NewMessage(chalk.Red, "That file doesn't exist!"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&justPrint, "print", "p", false, "Print out the script instead of running it")
	runCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Confirm you want to run the script")
}

func verifyYes(opts []wmenu.Opt) error {
	if opts[0].Value.(string) == "yes" {
		yes = true
	} else {
		yes = false
	}
	return nil
}
