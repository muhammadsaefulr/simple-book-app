package main

import (
	"github.com/muhammadsaefulr/simple-book-app/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
