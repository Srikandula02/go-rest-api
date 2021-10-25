package main

import (
	"go-api/config"
	"go-api/routers"
)

func main() {
	config.DBConnect()
	routers.SetupRouter()
}
