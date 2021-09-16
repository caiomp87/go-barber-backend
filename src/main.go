package main

import (
	"barber/src/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := routes.LoadUserRoutes()

	r.Run(":3333")
}
