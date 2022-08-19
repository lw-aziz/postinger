package main

import (
	"postinger/config"
	"postinger/library/mongodb"

	"postinger/util/localization"
	"postinger/util/logwrapper"

	"postinger/util/server"
	"strings"
)

func main() {
	configData := config.LoadEnv()

	Logger := logwrapper.NewLogger(configData.Server)
	localization.LoadBundle(configData.Server)

	Logger.Infoln(strings.Repeat("~", 50))
	mongoErr := mongodb.NewConnection(configData.MongoDB)
	if mongoErr != nil {
		Logger.Fatal("Error connecting to MongoDB : ", mongoErr)
	}
	server.StartServer(configData.Server)
}
