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
package main

import (
	"github.com/GreatGodApollo/qsr/cmd"
	"time"
)

func main() {
	cmd.Execute()
}

func test(){
	time.Sleep(3 * time.Second)
}