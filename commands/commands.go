package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func SetupApplicationCommand(app *cli.App) {
	setupBarkCommand(app)

	testCommand(app)
}

func setupBarkCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name: "bark",
		Usage:
		`send request,
example: (no config) 
	bark-cli [-X POST/GET] [-b=test-body] [-t=test-title] [-u=www.baidu.com] [-c=copy-content] [-a=true] bark
example: (config with flags)
	bark-cli [--host=https://api.day.app] [-p=443] [-k=xxxxxxxxxxx] [...other notification param] bark
example: (config with file)
	bark-cli [-f=/home/user/bark-cli/bark-cli.json] [...other notification param] bark`,
		Action: func(c *cli.Context) error {

		},
	})
}

func testCommand(app *cli.App) {
	// Create a Command
	app.Commands = append(app.Commands, &cli.Command{
		// 命令的名字
		Name: "test",
		// 命令的缩写，就是不输入language只输入lang也可以调用命令
		Aliases: []string{"t"},
		// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
		Usage: "test",
		// 命令的处理函数
		Action: func(c *cli.Context) error {
			fmt.Println("test")
			return nil
		},
	})
}
