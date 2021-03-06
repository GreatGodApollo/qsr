<img alt="Banner" src="assets/img/banner.png" style="border-radius:10px">
<h1 align="center">Quick Script Runner</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

Quick Script Runner (qsr) is a command line utility program that allows users to run code from github gists quickly with just a single command

## Installing

### Standard Download
Just head on over to the [releases](https://github.com/GreatGodApollo/qsr/releases) page and download the latest release
for your platform. Extract it using something like [7-Zip](https://www.7-zip.org) for Windows or `tar` on other 
platforms (`tar -zxvf qsr*.tar.gz`).

That's it! Although you'll probably want to also add the binary to your path for ease of use.

### Scoop
Do you happen to have [scoop](https://github.com/lukesampson/scoop) installed? Well lucky for you, I happen to have a [scoop](https://github.com/lukesampson/scoop) bucket. Said bucket's name is [Trough](https://github.com/GreatGodApollo/trough).
```shell
$ scoop bucket add trough https://github.com/GreatGodApollo/trough.git
$ scoop install qsr
```

### Go Get
Do you have go installed? You can run just one simple command to install QSR!
```shell
$ go get -u github.com/GreatGodApollo/qsr
```

## Usage

```shell
$ qsr --help
  Quick Script Runner is a command line utility that allows you to run gists
  with a single command.
  
  Usage:
    qsr [command]
  
  Available Commands:
    docs        Documentation Generator
    help        Help about any command
    link        Link a gist or file alias
    run         Run a remote gist
    source      Get a link to the source code
    unlink      Remove a gist or file alias
  
  Flags:
        --config string   config file (default is $HOME/.qsr.json)
    -h, --help            help for qsr
        --version         version for qsr
  
  Use "qsr [command] --help" for more information about a command.
```

### Configuration
These values can be set in `~/.qsr.json`. The options listed are unable to be used as alias names.

Configuration options available:

#### `token` - *Personal OAuth token for Github*
A personal OAuth token can be generated on your [tokens](https://github.com/settings/tokens) page. This allows for the
60 requests per hour limit to be surpassed.

***This token does not need any special scopes.***

## Included Gists
To prevent overcrowding this README, the included gist aliases are located in [gists.md](gists.md)

## Supported Languages
- Batch (Windows Only)
- Golang
- JavaScript
- Python 2 & 3 (Set with shebang ie. `#!/usr/bin/python3`)
- Ruby
- Shell

This list can be expanded on, either create an issue, or a PR to request a new language!

## Built With

* [wmenu](https://github.com/dixonwille/wmenu/)
* [go-cmd](https://github.com/go-cmd/cmd)
* [go-github](https://github.com/google/go-github)
* [Cobra](https://github.com/spf13/cobra)
* [Chalk](https://github.com/ttacon/chalk)


## Compiling

To compile the executable from the source, it's extremely easy, and can be done in as little as 3 commands.

You can follow these instructions to build:
```bash

# This assumes you already have git and golang installed.

$ git clone https://github.com/GreatGodApollo/qsr.git

$ cd qsr

$ go build

```

## Licensing

This project is licensed under the [GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)
