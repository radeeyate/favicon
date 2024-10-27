package main

import (
	"fmt"
	"io"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"

	"go.deanishe.net/favicon"
)

func main() {
	app := fiber.New()
	app.Use(cache.New())
	client := resty.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Get("/get/:domain", func(c *fiber.Ctx) error {
		domain := c.Params("domain")
		url := &url.URL{
			Scheme: "https",
			Host:   domain,
		}

		icons, err := favicon.Find(url.String())
		var resp *resty.Response
		var numErrors int
		for _, icon := range icons {
			if !icon.IsSquare() {
				continue
			}

			tempresp, err := client.R().SetDoNotParseResponse(true).Get(icon.URL)
			if tempresp.StatusCode() != 200 || err != nil {
				numErrors += 1
				continue
			} else {
				resp = tempresp
				break
			}
		}
		if err != nil || len(icons) == 0 {
			return c.SendFile("./default.ico") // we keep the status as 200 so we and the browser can cache it
		}

		if numErrors == len(icons) {
			return c.SendFile("./default.ico")
		}

		if resp != nil {
			defer resp.RawBody().Close()
		} else {
			return c.SendFile("./default.ico")
		}

		/* in the future we can resize the icons to reduce network
		throughput + loading times but it's not an issue as of now. */
		body, err := io.ReadAll(resp.RawBody())
		if err != nil {
			fmt.Println("read error:", err)
			c.Context().SetStatusCode(404)
			return c.SendFile("./default.ico")
		}

		contentType := resp.Header().Get("Content-Type")
		if contentType != "" {
			c.Set("Content-Type", contentType)
		} else {
			c.Set("Content-Type", "image/x-icon")
		}

		c.Context().SetStatusCode(200)
		return c.Send(body)
	})

	app.Listen(":7000")
}
