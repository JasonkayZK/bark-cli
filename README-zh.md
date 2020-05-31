<!-- PROJECT SHIELDS -->

 [![Build Status](https://travis-ci.org/JasonkayZK/bark-cli.svg?branch=master)](https://travis-ci.org/JasonkayZK/bark-cli) ![Go](https://img.shields.io/github/go-mod/go-version/JasonkayZK/bark-cli) ![repo-size](https://img.shields.io/github/repo-size/JasonkayZK/bark-cli) ![stars](https://img.shields.io/github/stars/JasonkayZK/bark-cli?style=social) ![MIT License](https://img.shields.io/github/license/JasonkayZK/bark-cli)

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
    <a href="https://github.com/JasonkayZK/bark-cli/issues">报告 Bug</a>
    ·
    <a href="https://github.com/JasonkayZK/bark-cli/pulls">提交新特性</a>
  </p>

</p>

<!-- TABLE OF CONTENTS -->

## 目录

* [关于本项目](#关于本项目)
* [用前须知](#用前须知)
    * [安装](#安装)
        -   [编译源代码安装](#编译源代码安装)
        -   [下载发行包](#下载发行包)
* [如何使用](#如何使用)
    -   [应用配置](#应用配置)
    -   [命令](#命令)
    -   [标志位](#标志位)
* [发行版本](#发行版本)
* [贡献代码](#贡献代码)
* [协议](#协议)
* [联系我](#联系我)



<!-- ABOUT THE PROJECT -->

## 关于本项目

[Bark](https://github.com/Finb/Bark)是一个简单但却很犀利的iOS应用，它允许你简单地发送Get/Post请求，将定制的通知推送到你的iPhone上。

这个项目允许您使用基本终端来使用Bark推送自定义通知！

像这样简单的命令即可推送通知：`bark-cli -t="Hello Title" -b="Hello body" bark`

由于这个项目是由Golang开发的，所以您可以在**任何平台**（Mac、Windows、Linux等）上使用这个应用程序

更多关于Bark的信息：

-   [Bark](https://github.com/Finb/Bark)
-   [bark-server](https://github.com/Finb/bark-server)

:star: 给本项目一个Star如果你觉得用得爽；

:fist: 通过Pull Request给本项目提供新的特性；

:question: 在Issue区提出你的问题；

<!-- GETTING STARTED -->

## 用前须知

### 安装

#### 编译源代码安装

您可以自己构建源代码。

执行此操作之前，请确保已安装`Golang>=v1.12`

##### 使用Go-Modules的方式安装

执行以下步骤来构建stardard go模块项目：

1.  Clone本仓库

```sh
git clone git@github.com:JasonkayZK/bark-cli.git
```

2.  编译源代码

```sh
go build
```

3.  安装源代码（这会将编译好的二进制文件复制到`SYSTEM $PATH`目录下）

```fsharp
go install
```

4.  测试安装

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

通过在终端中键入`bark cli`，若显示如上所述的帮助信息，说明安装成功:smile:

##### 通过Makefile编译

或者，你也可以使用“make”命令生成二进制文件：

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

二进制文件将在`bin/{platform}/bark-cli`下生成；

编译完成后还需要复制二进制文件到系统环境变量`$PATH`下，例如：

-   Win: ` C:\Windows\System32 `
-   Linux/Unix: `/usr/local/bin`

有关“make”命令的详细信息，请参见：[Makefile](https://github.com/JasonkayZK/bark-cli/blob/master/Makefile)

#### 下载发行包

直接下载对应平台的二进制文件：[Releases](https://github.com/JasonkayZK/bark-cli/releases)

然后将下载的二进制文件放在系统环境变量`$PATH`下，例如：

-   Win: ` C:\Windows\System32 `
-   Linux/Unix: `/usr/local/bin`

然后键入命令`bark cli`，来确保程序已正确安装：

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

看到上面的提示说明安装成功:smile:



<!-- USAGE EXAMPLES -->

## 如何使用

### 应用配置

#### 创建/更新配置

要使用Bark，您需要配置元数据，例如：主机、端口、密钥等

简单地说，可以使用`bark-cli config set`命令，生成或更新配置文件`$HOME/bark-cli/bark-cli.json`，例如：

```sh
$ bark-cli --host=your_host_address -p=8080 -k=xxxxxxxx config set
config updated at: C:\Users\Administrator\bark-cli\bark-cli.json
```

更简单的是，如果使用默认主机(`https://api.day.app`)和Bark提供的端口（`443`），则可以使用下面的命令生成配置文件：

```sh
$ bark-cli -k=123 config set
config updated at: C:\Users\Administrator\bark-cli\bark-cli.json
```

要更新设置文件，只需使用`bark-cli config set`，如下所示：

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

#### 列出设置

只需键入下面的命令即可列出在`$HOME/bark-cli/bark-cli.json`中的设置内容：

```sh
$ bark-cli config list
{
    "port": 8080,
    "host": "your_host_address",
    "key": "xxxxxxxx"
}
```

每次执行命令前将会首先加载此配置文件。

#### 配置优先级

您可以使用标志位来设置本次执行命令的配置，而不是使用默认的配置文件，例如：

```sh
$ bark-cli -p=port_num --host=host_address -k=xxxxx [other flags...] command
```

或者你也可以通过 `-f` or `--file`来指定通过其他配置文件传入配置信息

```sh
$ bark-cli -f=./bark-cli.json [other flags...] command
```

>**优先级： 指定文件 > 标志位 > 默认配置文件**
>
>注意，在执行命令之前将加载默认配置文件，然后指定的标志位会覆盖默认设置，然后`-f`覆盖标志位；
>
>所以当你执行类似于下面的命令时：
>
>```sh
>$ bark-cli -p=8888 --host=www.host.com -k=123333 -f=./bark-cli.json [other flags] command
>```
>
>实际的配置信息如下：
>
>```json
>{
>"port": 8888,
>"host": "www.host.com",
>"key": "123333"
>}
>```

### 命令

下面是本应用中的命令：

| 命令   | 解释                                                         | 发行日期   | 例子                                                         |
| ------ | ------------------------------------------------------------ | ---------- | ------------------------------------------------------------ |
| config | 创建/更新/显示 默认配置文件                                  | 2020-05-31 | `bark-cli -p=123 config set`<br />`bark-cli config list`     |
| bark   | 仅推送通知，携带标题(`-t/--titile`) 和消息体(`-b/--body`)    | 2020-05-31 | `bark-cli -t=bark-title -b=bark-body bark`                   |
| url    | 推送URL类型通知<br />点击通知可以打开推送的url               | 2020-05-31 | `bark-cli -t=url-title -b=url-body -u=https:www.baidu.com url` |
| copy   | 推送Copy类型通知<br />点击通知可复制copy信息，同时可以配置自动复制`-a=true` | 2020-05-31 | `bark-cli -t=copy-yourCode -b=code9527 -c=9527 -a=true copy` |

### 标志位

在执行命令时可以设置许多全局标志：

| 标志位        | 解释                                                     | 更新日期   |
| ------------- | -------------------------------------------------------- | ---------- |
| --host        | bark 服务器地址                                          | 2020-05-31 |
| --port/-p     | bark 服务器端口号                                        | 2020-05-31 |
| --key/-k      | bark 密钥<br />例如：`https://api.day.app/{key}/content` | 2020-05-31 |
| --title/-t    | 通知标题                                                 | 2020-05-31 |
| --body/-b     | 通知消息体                                               | 2020-05-31 |
| --barkUrl/-u  | URL通知                                                  | 2020-05-31 |
| --barkCopy/-c | Copy通知                                                 | 2020-05-31 |
| --autoCopy/-a | 启用自动复制功能<br />(默认：false)                      | 2020-05-31 |
| --request/-X  | 请求类型：GET or POST <br />(默认："POST")               | 2020-05-31 |
| --file/-f     | 从文件中读入配置信息                                     | 2020-05-31 |

通过 `bark-cli -h` 命令查看更多信息！

>   **几乎所有**标志都对应于：[Bark应用](https://github.com/Finb/Bark)



<!-- ROADMAP -->

## 发行版本

| 版本           | 发行状态 | 发行日期   | 备注           |
| -------------- | -------- | ---------- | -------------- |
| Release v1.0.0 | √        | 2020-05-31 | 实现了通用功能 |
| Release v1.0.1 | On-Going | -          | 增加定时通知   |
|                |          |            |                |



<!-- CONTRIBUTING -->

## 贡献代码

贡献使开源社区成为一个学习、激励和创造的绝佳场所。

你对本项目所作的任何贡献都是非常感谢的:heart:。

1. Fork本项目
2. 检出到新的分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送你的提交 (`git push origin feature/AmazingFeature`)
5. 开启一个新的Pull Request



<!-- LICENSE -->

## 协议

根据MIT许可协议发行。详细信息，请参见`LICENSE`。



<!-- CONTACT -->

## 联系我

Jasonkay - [Blog](https://jasonkayzk.github.io/) - 271226192@qq.com

项目地址: [https://github.com/JasonkayZK/bark-cli](https://github.com/JasonkayZK/bark-cli)

