package main

import (
	"dovran/mascot/config"
	"dovran/mascot/internal/app"
)

func main() {
	newConfig, err := config.NewConfig()
	if err != nil {
		return
	}
	app.Run(newConfig)
}
