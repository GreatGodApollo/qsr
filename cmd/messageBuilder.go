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

import "github.com/ttacon/chalk"

type Message struct {
	message string
}

func NewMessage(color chalk.Color, message string) *Message {
	return &Message{
		message: chalk.Cyan.Color("[QSR] ") + color.Color(message),
	}
}

func (msg *Message) ThenColor(color chalk.Color, message string) *Message {
	msg.message = msg.message + " " + color.Color(message)
	return msg
}

func (msg *Message) ThenStyle(style chalk.TextStyle, message string) *Message {
	msg.message = msg.message + " " + style.TextStyle(message)
	return msg
}

func (msg *Message) ThenColorStyle(color chalk.Color, style chalk.TextStyle, message string) *Message {
	msg.message = msg.message + " " + color.Color(style.TextStyle(message))
	return msg
}

func (msg *Message) String() string {
	return msg.message
}

