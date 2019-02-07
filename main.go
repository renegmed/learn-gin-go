package main

import "gin-go-mvansickle/app"

func main() {
	r := app.RegisterRoutes()

	r.Run(":3000")
}
