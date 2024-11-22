package main

import (
	"mailGOing/config"
	"mailGOing/internal/api"
)

func main() {
	config.LoadConfig()

	r := api.SetupRouter()
	r.Run(":8080")
}
