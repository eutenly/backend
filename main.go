package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Eutenly Web Services v0.1")

	//Create Echo Instance
	e := echo.New()

	//Setup API Router
	apiRouter(e)

	//Setup static router
	staticRouter(e)

	//Start server on 8080
	e.Logger.Fatal(e.Start(":8080"))
}
