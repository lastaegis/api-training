package main

import "github.com/labstack/echo/v4"

func main() {
	echoFramework := echo.New()
	RoutingAPI(echoFramework)
}
