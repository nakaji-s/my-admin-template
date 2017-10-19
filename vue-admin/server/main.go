package main

import (
	"github.com/labstack/echo"
)

func main() {
	// start server
	e := echo.New()

	e.Static("/", "dist")

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
