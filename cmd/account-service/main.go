package main

import (
	"account-service/internal/app"
	"log"
)

func main() {
	c := app.NewContext()
	log.Fatal(c.Server().Start())

}
