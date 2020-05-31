package flags

import (
	"github.com/jasonkayzk/bark-cli/utils"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

const (
	HostValueDefault = "https://api.day.app"
	PortValueDefault = 443
	KeyValueDefault  = ""
)

var DefaultConfigPath string

var hostValue string
var portValue int64
var keyValue string

// init the config flags: host, port & key from DefaultConfigPath or Default value
func init() {
	DefaultConfigPath = utils.Home() + string(filepath.Separator) + "bark-cli" + string(filepath.Separator) + "bark-cli.json"
	if utils.ConfigExist(DefaultConfigPath) {
		config, err := utils.LoadConfig(DefaultConfigPath)
		if err != nil {
			hostValue = HostValueDefault
			portValue = PortValueDefault
			keyValue = KeyValueDefault
			return
		}

		if config.Host != "" {
			hostValue = config.Host
		} else {
			hostValue = HostValueDefault
		}
		if config.Port > 0 {
			portValue = config.Port
		}else {
			portValue = PortValueDefault
		}
		if config.Key != "" {
			keyValue = config.Key
		} else {
			keyValue = KeyValueDefault
		}
	} else {
		hostValue = HostValueDefault
		portValue = PortValueDefault
		keyValue = KeyValueDefault
	}
}

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
		Name:    "barkUrl",
		Aliases: []string{"u"},
		Value:   "",
		Usage:   "notification url",
	})
}

func setCopyContentFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "barkCopy",
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
		Value:   portValue,
		Usage:   "bark server port number",
	})
}

func setHostFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:  "host",
		Value: hostValue,
		Usage: "bark server host location",
	})
}

func setKeyFlag(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Value:   keyValue,
		Usage:   "secret key from bark, such as: https://api.day.app/{key}/content",
	})
}


