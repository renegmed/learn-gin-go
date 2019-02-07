package main

import "gin-go-mvansickle/app"

func main() {
	r := app.RegisterRoutes()
	r.Static("/public", "./public")
	r.Run(":3000")
}
