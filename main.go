package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		filetype := file.Header.Get("Content-Type")
		if filetype != "image/jpeg" && filetype != "image/png" {
			return fmt.Errorf("file type not supported")
		}
		name := "img"
		time_unix := time.Now().Unix()
		time_stamp := strconv.Itoa(int(time_unix))
		ext := filepath.Ext(file.Filename)
		name_file := fmt.Sprintf("./upload/%s-%s%s", name, time_stamp, ext)
		// name_file =
		c.SaveFile(file, name_file)
		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
		})

	})

	log.Fatal(app.Listen(":3000"))
}
