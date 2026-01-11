package main

import (
	"morhaat.com/mh-ui-service/src/config"
	"morhaat.com/mh-ui-service/src/utils"

	"morhaat.com/mh-ui-service/src/router"
)

func main() {
	utils.InitLogger()
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.Log.Fatalf("Failed to load configuration: %v", err)
	}
	utils.Log.Infof("Configuration loaded: %+v", cfg)
	router.ConfigureRouter()
}
