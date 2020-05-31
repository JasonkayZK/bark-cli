package commands

import (
	"bytes"
	"fmt"
	"github.com/jasonkayzk/bark-cli/utils"
	"log"
	"strings"

	"github.com/urfave/cli/v2"
)

func SetupApplicationCommand(app *cli.App) {
	setupBarkCommand(app)
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
			var host string
			var port int64
			var key string
			// Priority: file > cli > default;

			// Step 1: Get url param
			// Step 1.1: has cli param
			host = c.String("host")
			port = c.Int64("port")
			key = c.String("key")

			// Step 1.2: has cliFile param
			cliFile := c.String("file")
			if cliFile != "" {
				config, err := utils.LoadConfig(cliFile)
				if err != nil {
					return err
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
					return fmt.Errorf("invalid args: [host, port, key]")
				}
			}

			// Step 1.3: else default cli param...

			// Step 2: send request, default requestType: POST
			// Step 2.1: GET request
			if strings.ToUpper(c.String("request")) == "GET" {
				getUrl, err := generateGetUrl(c)
				if err != nil {
					return fmt.Errorf("GET request failed, err :%s", err)
				}

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
			postUrl, err := generatePostUrl(c)
			if err != nil {
				return fmt.Errorf("generate post url failed, err :%s", err)
			}

			post, err := utils.Post(postUrl, map[string]string{
				"title": c.String("title"),
				"body":  c.String("body"),
				"url":   c.String("url"),
				"copy":  c.String("copy"),
				"automaticallyCopy": string((func(b bool) int64 {
					if b {
						return 1
					}
					return 0
				})(c.Bool("autoCopy"))),
			})
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

func generateGetUrl(c *cli.Context) (string, error) {
	buffer := bytes.Buffer{}

	// url prefix
	buffer.WriteString(fmt.Sprintf("%s:%s/%s/%s/?", c.String("host"), string(c.Int64("port")), c.String("key"), c.String("body")))

	// url query param




}

func generatePostUrl(c *cli.Context) (string, error) {

}
