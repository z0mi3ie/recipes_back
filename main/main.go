package main

import (
	"github.com/z0mi3ie/recipes_back/server"
)

func main() {
	server := server.New()
	server.Start()
}
