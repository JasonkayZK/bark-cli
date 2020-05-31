package flags

import (
	"github.com/urfave/cli/v2"
)

func SetupApplicationFlags(app *cli.App) {
	setHostFlag(app)
	setPortFlag(app)
	setKeyFlag(app)
	setTitleFlag(app)
	setBodyFlag(app)
	setUrlFlag(app)
	setCopyContentFlag(app)
	setAutoCopyFlag(app)
	setRequestTypeFlag(app)
	setConfigFileFlag(app)
}

func setTitleFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "title",
		Aliases: []string{"t"},
		Value:   "",
		Usage:   "notification title",
	})
}

func setBodyFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "body",
		Aliases: []string{"b"},
		Value:   "Empty Body",
		Usage:   "notification body",
	})
}

func setUrlFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "",
		Usage:   "notification url",
	})
}

func setCopyContentFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "copy",
		Aliases: []string{"c"},
		Value:   "",
		Usage:   "notification copy content",
	})
}

func setAutoCopyFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.BoolFlag{
		Name:    "autoCopy",
		Aliases: []string{"a"},
		Value:   false,
		Usage:   "enable automaticallyCopy for notification",
	})
}

func setRequestTypeFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "request",
		Aliases: []string{"X"},
		Value:   "POST",
		Usage:   "request method: GET or POST",
	})
}

func setConfigFileFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Value:   "",
		Usage:   "config bark-cli parameter from json file, such as: host, port, key, ...",
	})
}

func setPortFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.Int64Flag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   443,
		Usage:   "bark server port number",
	})
}

func setHostFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:  "host",
		Value: "https://api.day.app",
		Usage: "bark server host location",
	})
}

func setKeyFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Value:   "",
		Usage:   "secret key from bark, such as: https://api.day.app/{key}/content",
	})
}
