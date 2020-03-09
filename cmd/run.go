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
	"github.com/ttacon/chalk"

	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var justPrint bool
var yes       bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [gist id] [filename] [arguments]",
	Short: "Run a remote gist",
	Long: `Quick Script Runner is a command line utility that allows you to run gists
with a single command.`,
	Args:  cobra.MinimumNArgs(2),
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
		file := gist.Files[github.GistFilename(args[1])	]
		if justPrint {
			fmt.Println(chalk.Cyan.Color("[QSR]"), "FileType:",chalk.Blue.Color(file.GetLanguage()))
			a := strings.Split(file.GetContent(), "\n")
			b := strings.Join(a, "\n" + chalk.Cyan.Color("[QSR] "))
			fmt.Println(chalk.Cyan.Color("[QSR]"), b)
		} else {

			if file.GetSize() > 0 {
				if !yes {
					menu := wmenu.NewMenu(chalk.Cyan.Color("[QSR] ") + chalk.Red.Color("Are you sure you want to run this script?"))
					menu.IsYesNo(wmenu.DefN)
					menu.Action(verifyYes)
					err := menu.Run()
					if err != nil {
						fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color(err.Error()))
						return
					}
					if !yes {return}
				}
				startingDirectory, err := os.Getwd()
				if CheckError(err) {
					return
				}

				if runtime.GOOS == "windows" {
					err = os.Chdir(os.Getenv("TEMP"))
					if CheckError(err) {return}
				} else if runtime.GOOS == "linux" {
					if _, err := os.Stat("/tmp/qsr/"); os.IsNotExist(err) {
						err = os.Mkdir("/tmp/qsr/", os.ModeDir)
						if CheckError(err) {return}
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
							if CheckError(err) {return}
							err = RunCommand("tmp.bat", args[2:]...)
							CheckError(err)
							err = os.Remove("tmp.bat")
							if CheckError(err) {return}
							break
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't " +
								"supported on your system!"))
							return
						}
					}
				case "Shell":
					{
						if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
							err = DownloadFile("tmp.sh", file.GetRawURL())
							if CheckError(err) {return}
							err = RunCommand("chmod", "+x", "tmp.sh")
							CheckError(err)
							err = RunCommand("./tmp.sh", args[2:]...)
							CheckError(err)
							err = os.Remove("tmp.sh")
							if CheckError(err) {return}
							break
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't " +
								"supported on your system!"))
							return
						}
					}
				case "Go":
					{
						err = DownloadFile("tmp.go", file.GetRawURL())
						if CheckError(err) {return}
						a := []string{"run", "tmp.go"}
						a = append(a, args[2:]...)
						err = RunCommand("go", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("You can find instructions " +
								"to install it here:"), chalk.Red.Color(chalk.Underline.TextStyle("https://golang.org")))
						}
						err = os.Remove("tmp.go")
						if CheckError(err) {return}
						break
					}
				case "JavaScript":
					{

						err = DownloadFile("tmp.js", file.GetRawURL())
						if CheckError(err) {return}
						a := []string{"tmp.js"}
						a = append(a, args[2:]...)
						err = RunCommand("node", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("You can find instructions " +
								"to install it here:"), chalk.Red.Color(chalk.Underline.TextStyle("https://nodejs.org/")))
						}
						err = os.Remove("tmp.js")
						if CheckError(err) {return}
						break
					}
				case "Python":
					{
						err = DownloadFile("tmp.py", file.GetRawURL())
						if CheckError(err) {return}
						fst, err := getFirstLine("tmp.py")
						if CheckError(err) {return}
						if fst == "#!/usr/bin/python3" {
							a := []string{"tmp.py"}
							a = append(a, args[2:]...)
							err = RunCommand("python3", a...)
							CheckError(err)
							if err != nil && strings.Contains(err.Error(), "executable file not found") {
								fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("You can find instructions " +
									"to install it here:"), chalk.Red.Color(chalk.Underline.TextStyle("https://www.python.org")))
							}
						} else if fst == "#!/usr/bin/python" {
							a := []string{"tmp.py"}
							a = append(a, args[2:]...)
							err = RunCommand("python", a...)
							CheckError(err)
							if err != nil && strings.Contains(err.Error(), "executable file not found") {
								fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("You can find instructions " +
									"to install it here:"), chalk.Red.Color(chalk.Underline.TextStyle("https://www.python.org")))
							}
						} else {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("Cannot determine language!"))
						}
						err = os.Remove("tmp.py")
						if CheckError(err) {return}
						break
					}
				case "Ruby":
					{
						err = DownloadFile("tmp.rb", file.GetRawURL())
						if CheckError(err) {return}
						a := []string{"tmp.rb"}
						a = append(a, args[2:]...)
						err = RunCommand("ruby", a...)
						CheckError(err)
						if err != nil && strings.Contains(err.Error(), "executable file not found") {
							fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("You can find instructions " +
								"to install it here:"), chalk.Red.Color(chalk.Underline.TextStyle("https://ruby-lang.org/")))
						}
						err = os.Remove("tmp.rb")
						if CheckError(err) {return}
						break
					}
				default:
					{
						fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't supported!"))
						break
					}
				}
				err = os.Chdir(startingDirectory)
				if CheckError(err) {return}
			} else {
				fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That file doesn't exist!"))
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