package main

import (
	"fmt"
	"os"

	"cth.release/go-rds/common"
	"cth.release/go-rds/web"
)

func main() {
	port := common.ThreeTermString(len(os.Getenv("PORT")) > 0, os.Getenv("PORT"), "8080")

	app := web.InitServer()

	app.App.Listen(fmt.Sprint(":" + port))
}
