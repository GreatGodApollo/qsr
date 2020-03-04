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
	"bufio"
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/ttacon/chalk"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func CheckError(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "executable file not found") {
			fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color("That language isn't supported on your system!"))
			return true
		}
		fmt.Println(chalk.Cyan.Color("[QSR]"), chalk.Red.Color(err.Error()))
		return true
	}
	return false
}

func RunCommand(command string, args ...string) error {
	// Create Cmd with options
	cmO := cmd.Options{
		Buffered: false,
		Streaming: true,
	}
	envCmd := cmd.NewCmdOptions(cmO, command, args[0], args[1])

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	stat := <-envCmd.Start()
	// Wait for goroutine to print everything
	<-doneChan
	return stat.Error
}

func getFirstLine(path string) (string, error) {

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	snl.Scan()
	txt := snl.Text()
	if err := snl.Err(); err != nil {return "", err}
	return txt, nil
}