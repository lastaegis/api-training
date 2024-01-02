package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	// Start echo
	echoFramework := echo.New()
	RoutingAPI(echoFramework)
}
