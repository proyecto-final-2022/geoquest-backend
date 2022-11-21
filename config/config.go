package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	APP_NOTIFICATIONS_URL string
}

func GetConfig(env string) Configuration {
	configuration := Configuration{}
	//	if len(params) > 0 {
	//		env = params[0]
	//	}
	fileName := fmt.Sprintf("./prod_config.json")
	gonfig.GetConf(fileName, &configuration)
	fmt.Println("***Conf: ", configuration.APP_NOTIFICATIONS_URL)
	return configuration
}
