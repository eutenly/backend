package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	fmt.Println("Eutenly Web Services v0.1")

	//Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Create Echo Instance
	e := echo.New()

	//Setup API Router
	apiRouter(e)

	//Setup static router
	staticRouter(e)

	//Enable Sessions
	sessionManager := session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET"))))
	e.Use(sessionManager)

	//Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
