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

import "github.com/spf13/viper"

func SetGists() {
	/* Repository of different user's gists!
	 * If you'd like your personal gist here, create a pull request, adding it to this list,
	 * and to gists.md
	 */


	//// "apollo": GreatGodApollo (Author)
	viper.Set("apollo", map[string]string{"gist": "00d91bfb540b8bc0169606b4d4f740a3"})

	//// "324luke": 324Luke (Contributor)
	viper.Set("324luke", map[string]string{"gist": "3a30e079e8692cc09776186014090835"})
	
	//// "Hello world": 324Luke (Contributor)
	viper.Set("hello-world", map[string]string{"gist": "6240b6a181305b5f5ac7a70763703448"})
}
