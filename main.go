package main

import (
	"github.com/lef237/gin-mcs/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // http://localhost:8080
}
