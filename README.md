<img alt="Banner" src="assets/img/banner.png" style="border-radius:10px">
<h1 align="center">Quick Script Runner</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

Quick Script Runner (qsr) is a command line utility program that allows users to run code from github gists quickly with just a single command

## Built With

* [wmenu](https://github.com/dixonwille/wmenu/)
* [go-cmd](https://github.com/go-cmd/cmd)
* [go-github](https://github.com/google/go-github)
* [Cobra](github.com/spf13/cobra)
* [Chalk](github.com/ttacon/chalk)

## Install

At the moment, you'll have to manually build the executable from the source, although this shouldn't be hard.

You can follow these instructions to build:
```bash

# This assumes you already have git and golang installed.

$ git clone https://github.com/GreatGodApollo/qsr.git

$ cd qsr

$ go build

```

That's it! Although you'll probably want to also add the binary to your path for ease of use.

## Usage

```bash
$ qsr --help

Quick Script Runner is a command line utility that allows you to run gists
with a single command.

Usage:
  qsr [command]

Available Commands:
  docs        Documentation Generator
  help        Help about any command
  run         Run a remote gist

Flags:
  -h, --help      help for qsr
      --version   version for qsr

Use "qsr [command] --help" for more information about a command.
```

## Licensing

This project is licensed under the [GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)
