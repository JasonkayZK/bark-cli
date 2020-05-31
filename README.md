<!-- PROJECT SHIELDS -->

 [![Build Status](https://travis-ci.org/JasonkayZK/bark-cli.svg?branch=master)](https://travis-ci.org/JasonkayZK/bark-cli) ![Go](https://img.shields.io/github/go-mod/go-version/JasonkayZK/bark-cli) ![repo-size](https://img.shields.io/github/repo-size/JasonkayZK/bark-cli)![stars](https://img.shields.io/github/stars/JasonkayZK/bark-cli?style=social) ![MIT License](https://img.shields.io/github/license/JasonkayZK/bark-cli)

<!-- PROJECT LOGO -->

<br />

<p align="center">
  <a href="https://github.com/JasonkayZK/bark-cli">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Bark-Cli</h3>
  <p align="center">
    A simple terminal tool for <a href="https://github.com/Finb/Bark">bark</a>
    <br />
    <br />
    <a href="https://github.com/JasonkayZK/bark-cli/README-zh.md">简体中文</a> ·
    <a href="https://github.com/JasonkayZK/bark-cli/README.md">English</a>·
    <a href="https://github.com/JasonkayZK/bark-cli/issues">Report Bug</a>
    ·
    <a href="https://github.com/JasonkayZK/bark-cli/pulls">Request Feature</a>
  </p>

</p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

* [About the Project](#about-the-project)
* [Getting Started](#getting-started)
  * [Installation](#installation)
      -   [Build Source Code](#build-source-code)
      -   [Download Release](#download-release)
* [Usage](#usage)
    -   [Settings](#settings)
    -   [Commands](#commands)
    -   [Flags](#flags)
* [Release](#releases)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->

## About The Project

[Bark](https://github.com/Finb/Bark) is a simple but sharp iOS App which allows you to simply send Get/Post request to push customed notifications to your iPhone.

This project allows you to use basicly terminal to use bark to push customed notifications!

Simple Like this: `bark-cli -t="Hello-Title" -b="Hello-body" bark`

And notice this project is developed by Golang, so you can use this application for **any platform**(Mac, Windows, Linux and so on)

More infomation about Bark:

-   [Bark](https://github.com/Finb/Bark)
-   [bark-server](https://github.com/Finb/bark-server)

:star: Star this project if you enjoy this application;

:fist: Pull requests for better ideas;

:question: Leave a issue for more questions;

<!-- GETTING STARTED -->

## Getting Started

### Installation

#### Build Source Code

You can build Source code by yourself.

Before doing so, make sure that you have installed `Golang >= v1.12`.

##### Install as Go-Modules project

Then do following steps to build a stardard go-modules project

2. Clone this repository
```sh
git clone git@github.com:JasonkayZK/bark-cli.git
```
2.  Build the source code

```sh
go build
```
3.  Install the source code(this with copy the binary file to your `SYSTEM $PATH`)

```fsharp
go install
```

4.  Test Installation

```sh
$ bark-cli
NAME:
   bark-cli - A simple cli for bark

USAGE:
   bark-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0
......
```

By type `bark-cli` in your terminal, its will show the helper as above;

Then, your installation is finished:smile:

##### Build by Makefile

Or, Your can Use `make` command to build the binary file:

```sh
# build all platform binary file, include mac, linux and win
$ make 
or
$ make all

# build mac only
$ make build-mac

# build linux only
$ make build-linux

# build windows only
$ make build-win
```

the binary file will be generated under `bin/{platform}/bark-cli`

Your may need to copy the binary file under your `SYSTEM $PATH`, such as:

-   Win: ` C:\Windows\System32 `
-   Linux/Unix: `/usr/local/bin`

For more information about `make` command, see: [Makefile](https://github.com/JasonkayZK/bark-cli/blob/master/Makefile)

#### Download Release

Download the corresponding built file at [Releases](https://github.com/JasonkayZK/bark-cli/releases)

Then Copy the binary file under your `SYSTEM $PATH`, such as:

-   Win: ` C:\Windows\System32 `
-   Linux/Unix: `/usr/local/bin`

Then type the command `bark-cli`, to make sure the program is installed:

```sh
$ bark-cli
NAME:
   bark-cli - A simple cli for bark

USAGE:
   bark-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0
......
```

Your installation is done after seen information above:smile:



<!-- USAGE EXAMPLES -->

## Usage

### Settings

#### Generate/Update Settings

To use bark, your need to config the meta data such as: host, port, key, ...

Simply, you can use `bark-cli config set` to generate or update a setting file at `$HOME/bark-cli/bark-cli/json`, such as:

```sh
$ bark-cli --host=your_host_address -p=8080 -k=xxxxxxxx config set
config updated at: C:\Users\Administrator\bark-cli\bark-cli.json
```

More simple, if you use the default host(`https://api.day.app`) and port(`443`) as Bark provided, you can generate your config file like this:

```sh
$ bark-cli -k=123 config set
config updated at: C:\Users\Administrator\bark-cli\bark-cli.json
```

To update the setting file, just use `bark-cli config set` as below:

```sh
$ bark-cli config list
{
    "port": 8080,
    "host": "your_host_address",
    "key": "xxxxxxxx"
}

$ bark-cli -p=123 config set
config updated at: C:\Users\Administrator\bark-cli\bark-cli.json

$ bark-cli config list
{
    "port": 123,
    "host": "your_host_address",
    "key": "xxxxxxxx"
}

```

#### Show Settings

Simply you can list your settings contents at `$HOME/bark-cli/bark-cli/json` by typing:

```sh
$ bark-cli config list
{
    "port": 8080,
    "host": "your_host_address",
    "key": "xxxxxxxx"
}
```

This setting file will be loaded when you execute a command.

#### Priority

Instead of using default setting file, you can use flags to set the settings in current execution, such as:

```sh
$ bark-cli -p=port_num --host=host_address -k=xxxxx [other flags...] command
```

Even you can load the settings from other setting files by using `-f` or `--file`:

```sh
$ bark-cli -f=./bark-cli.json [other flags...] command
```

>**Priority: file > flags > default**
>
>Notice that default config file will be load before execute the command, then flags override the default settings, and `-f` param override the flags in the end;
>
>So, if you execute like this:
>
>```sh
>$ bark-cli -p=8888 --host=www.host.com -k=123333 -f=./bark-cli.json [other flags] command
>```
>
>The setting is always:
>
>```json
>{
>    "port": 8888,
>    "host": "www.host.com",
>    "key": "123333"
>}
>```

### Commands

All commands is in the list below:

| Command | Explain                                                      | Release Date | Example                                                      |
| ------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| config  | generate/update/list the default config                      | 2020-05-31   | `bark-cli -p=123 config set`<br />`bark-cli config list`     |
| bark    | Push the notifications with title(`-t/--titile`) and body(`-b/--body`) | 2020-05-31   | `bark-cli -t=bark-title -b=bark-body bark`                   |
| url     | Push the url notification.<br />You can open the url by click the notification. | 2020-05-31   | `bark-cli -t=url-title -b=url-body -u=https:www.baidu.com url` |
| copy    | Push the copy notification.<br />Your can copy or enable auto-copy | 2020-05-31   | `bark-cli -t=copy-yourCode -b=code9527 -c=9527 -a=true copy` |

### Flags

There are plenty global flags you can set in execution:

| Flag          | Explain                                                      | Update Date |
| ------------- | ------------------------------------------------------------ | ----------- |
| --host        | bark server host location                                    | 2020-05-31  |
| --port/-p     | bark server port number                                      | 2020-05-31  |
| --key/-k      | secret key from bark<br />such as: `https://api.day.app/{key}/content` | 2020-05-31  |
| --title/-t    | notification title                                           | 2020-05-31  |
| --body/-b     | notification body                                            | 2020-05-31  |
| --barkUrl/-u  | notification url                                             | 2020-05-31  |
| --barkCopy/-c | notification copy content                                    | 2020-05-31  |
| --autoCopy/-a | enable automaticallyCopy for notification<br />(default: false) | 2020-05-31  |
| --request/-X  | request method: GET or POST <br />(default: "POST")          | 2020-05-31  |
| --file/-f     | config bark-cli parameter from json file                     | 2020-05-31  |

For more information, just type `bark-cli -h` for help!

>   **Almost all** flags are corresponding to the origin [Bark Application](https://github.com/Finb/Bark)



<!-- ROADMAP -->

## Releases

| Version        | Status   | Date       | Note           |
| -------------- | -------- | ---------- | -------------- |
| Release v1.0.0 | √        | 2020-05-31 | For common use |
| Release v1.0.1 | On-Going | -          | Add Cron       |
|                |          |            |                |



<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->

## Contact

Jasonkay - [Blog](https://jasonkayzk.github.io/) - 271226192@qq.com

Project Link: [https://github.com/JasonkayZK/bark-cli](https://github.com/JasonkayZK/bark-cli)

