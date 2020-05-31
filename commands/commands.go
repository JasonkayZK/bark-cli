package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jasonkayzk/bark-cli/flags"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/jasonkayzk/bark-cli/utils"

	"github.com/urfave/cli/v2"
)

const (
	barkCommandUsage = `Send bark request
	example: (default config) 
		bark-cli [-X POST/GET] [-b=bark-body] [-t=bark-title] bark
	example: (config with flags)
		bark-cli [-X POST/GET] [--host=https://api.day.app] [-p=443] [-k=xxxxxxxxxxx] [...other notification param] bark
	example: (config with file)
		bark-cli [-f=/home/user/bark-cli/bark-cli.json] [...other notification param] bark
`

	barkUrlUsage = `Send barkUrl request
	example: (default config) 
		bark-cli [-X POST/GET] [-t=urlBaidu-title] [-b=test-baiduUrl] -u=https://www.baidu.com url
	example: (config with flags)
		bark-cli [-X POST/GET] [--host=https://api.day.app] [-p=443] [-k=xxxxxxxxxxx] [...other notification param] url
	example: (config with file)
		bark-cli [-f=/home/user/bark-cli/bark-cli.json] [...other notification param] url
`

	barkCopyUsage = `Send barkCopy request
	example: (default config) 
		bark-cli [-X POST/GET] [-t=copy-yourCode] [-b=code9527] -c=9527 -a=true copy
	example: (config with flags)
		bark-cli [-X POST/GET] [--host=https://api.day.app] [-p=443] [-k=xxxxxxxxxxx] [...other notification param] copy
	example: (config with file)
		bark-cli [-f=/home/user/bark-cli/bark-cli.json] [...other notification param] copy
`

	configUsage = `Settings about bark-cli
	example: (set config)
		bark-cli [--host=https://api.day.app] [-p=443] -k=xxxxxxxxxxx config set
	example: (list config)
		bark-cli config list
`

	configSetUsage = `Generator(Override) bark-cli config file at $HOME/bark-cli/bark-cli.json
	example: 
		bark-cli [--host=https://api.day.app] [-p=443] -k=xxxxxxxxxxx config set
`

	configListUsage = `List bark-cli config
	example:
		back-cli config list
`
)

func SetupApplicationCommand(app *cli.App) {
	setBarkNotificationCommand(app)
	setUrlNotificationCommand(app)
	setCopyNotificationCommand(app)
	setGenerateConfigCommand(app)
}

func setBarkNotificationCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "bark",
		Usage: barkCommandUsage,
		Action: func(c *cli.Context) error {
			// Step 1: Get config param
			// Priority: file > cli > default;
			host, port, key, err := getConfig(c)
			if err != nil {
				return err
			}

			// Step 2: send request, default requestType: POST
			// Step 2.1: GET request
			if strings.ToUpper(c.String("request")) == "GET" {
				getUrl := generateGetUrlPrefix(host, port, key, c)

				log.Printf("get url: %s", getUrl)

				get, err := utils.Get(getUrl)
				if err != nil {
					return fmt.Errorf("generate get url failed, err :%s", err)
				}
				if get.Code != 200 {
					return fmt.Errorf("get request failed, err :%s", err)
				}

				log.Println("Send notification success!")
				return nil
			}

			// Step 2.2: POST request as default
			postUrl := generatePostUrlPrefix(host, port, key)

			post, err := utils.Post(postUrl, &url.Values{
				"title": []string{c.String("title")},
				"body":  []string{c.String("body")},
			})
			log.Printf("Post url: %s\n", postUrl)
			if err != nil {
				return fmt.Errorf("post request failed")
			}
			if post.Code != 200 {
				return fmt.Errorf("post request failed, err :%s", err)
			}

			log.Println("Send notification success!")
			return nil
		},
	})
}

func setUrlNotificationCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "url",
		Usage: barkUrlUsage,
		Action: func(c *cli.Context) error {
			host, port, key, err := getConfig(c)
			if err != nil {
				return err
			}

			if strings.ToUpper(c.String("request")) == "GET" {
				getUrl := generateGetUrlPrefix(host, port, key, c)

				getUrl += fmt.Sprintf("?url=%s/", c.String("barkUrl"))

				log.Printf("get url: %s", getUrl)

				get, err := utils.Get(getUrl)
				if err != nil {
					return fmt.Errorf("generate get url failed, err :%s", err)
				}
				if get.Code != 200 {
					return fmt.Errorf("get request failed, err :%s", err)
				}

				log.Println("Send notification success!")
				return nil
			}

			// Step 2.2: POST request as default
			postUrl := generatePostUrlPrefix(host, port, key)

			log.Println()

			post, err := utils.Post(postUrl, &url.Values{
				"title": []string{c.String("title")},
				"body":  []string{c.String("body")},
				"url":   []string{c.String("barkUrl")},
			})
			log.Printf("Post url: %s\n", postUrl)
			if err != nil {
				return fmt.Errorf("post request failed")
			}
			if post.Code != 200 {
				return fmt.Errorf("post request failed, err :%s", err)
			}

			log.Println("Send notification success!")
			return nil
		},
	})
}

func setCopyNotificationCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "copy",
		Usage: barkCopyUsage,
		Action: func(c *cli.Context) error {
			host, port, key, err := getConfig(c)
			if err != nil {
				return err
			}

			// Step 2: send request, default requestType: POST
			// Step 2.1: GET request
			if strings.ToUpper(c.String("request")) == "GET" {
				getUrl := generateGetUrlPrefix(host, port, key, c)

				log.Printf("barkCopy: %s\n", c.String("barkCopy"))

				getUrl += fmt.Sprintf("?copy=%s", c.String("barkCopy"))
				if c.Bool("autoCopy") {
					getUrl += "&automaticallyCopy=1"
				}

				log.Printf("get url: %s", getUrl)

				get, err := utils.Get(getUrl)
				if err != nil {
					return fmt.Errorf("generate get url failed, err :%s", err)
				}
				if get.Code != 200 {
					return fmt.Errorf("get request failed, err :%s", err)
				}

				log.Println("Send notification success!")
				return nil
			}

			// Step 2.2: POST request as default
			postUrl := generatePostUrlPrefix(host, port, key)

			post, err := utils.Post(postUrl, &url.Values{
				"title": []string{c.String("title")},
				"body":  []string{c.String("body")},
				"copy":  []string{c.String("barkCopy")},
				"automaticallyCopy": []string{string((func(b bool) int64 {
					if b {
						return 1
					}
					return 0
				})(c.Bool("autoCopy")))},
			})
			log.Printf("Post url: %s\n", postUrl)
			if err != nil {
				return fmt.Errorf("post request failed")
			}
			if post.Code != 200 {
				return fmt.Errorf("post request failed, err :%s", err)
			}

			log.Println("Send notification success!")
			return nil
		},
	})
}

func setGenerateConfigCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "config",
		Usage: configUsage,
		Subcommands: []*cli.Command{
			{
				Name:  "set",
				Usage: configSetUsage,
				Action: func(c *cli.Context) error {
					host, port, key := c.String("host"), c.Int64("port"), c.String("key")
					if host == "" {
						host = flags.HostValueDefault
					}
					if port <= 0 {
						port = flags.PortValueDefault
					}
					if key == "" {
						key = flags.KeyValueDefault
					}

					ok, err := utils.ReplaceConfig(flags.DefaultConfigPath, utils.ConfigParam{
						Host: host,
						Port: port,
						Key:  key,
					})
					if err != nil || !ok {
						return fmt.Errorf("write config err: %s\n", err)
					}
					return nil
				},
			},
			{
				Name:  "list",
				Usage: configListUsage,
				Action: func(c *cli.Context) error {
					conf, err := utils.LoadConfig(flags.DefaultConfigPath)
					if err != nil {
						return fmt.Errorf("load config err: %s", err)
					}

					str, err := json.MarshalIndent(conf, "", "    ")
					if err != nil {
						return fmt.Errorf("marshal err: %s", err)
					}

					fmt.Println(string(str))

					return nil
				},
			},
		},
	})
}

// Get config param
// Priority: file > cli > default;
func getConfig(c *cli.Context) (string, int64, string, error) {
	// Step 1: Get url param
	// Step 1.1: has cli param
	host, port, key := c.String("host"), c.Int64("port"), c.String("key")

	// Step 1.2: has cliFile param
	cliFile := c.String("file")
	if cliFile != "" {
		config, err := utils.LoadConfig(cliFile)
		if err != nil {
			return "", 0, "", err
		}

		if config.Host != "" {
			host = config.Host
		}
		if config.Port > 0 {
			port = config.Port
		}
		if config.Key != "" {
			key = config.Key
		}

		if host == "" || port <= 0 || key == "" {
			return "", 0, "", fmt.Errorf("invalid args: [host, port, key]")
		}
	}

	// Step 1.3: else default cli param...
	return host, port, key, nil
}

func generateGetUrlPrefix(host string, port int64, key string, c *cli.Context) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%s:%s/%s/%s/%s", host, strconv.FormatInt(port, 10), key, c.String("title"), c.String("body")))
	return buffer.String()
}

func generatePostUrlPrefix(host string, port int64, key string) string {
	return fmt.Sprintf("%s:%s/%s/", host, strconv.FormatInt(port, 10), key)
}
