package server

import (
	"fmt"
	"postinger/config"
	"postinger/util/logwrapper"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// StartServer - start Server server
func StartServer(serverConfig config.ServerConfig) {

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.New()

	/*
		Define middleware
	*/
	router.Use(setUserStatus())
	// set logger in gin
	router.Use(logwrapper.GinLogger(), gin.Recovery())

	// Initialize the routes
	initializeRoutes()

	// middlewareFunc := cors.New(cors.Options{
	// 	AllowedOrigins:   serverConfig.AllowOrigins,
	// 	AllowedMethods:   []string{"POST", "GET", "DELETE", "PATCH", "PUT"},
	// 	AllowedHeaders:   []string{"Origin", "Authorization"},
	// 	ExposedHeaders:   []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           int(time.Duration(12 * time.Hour).Seconds()),
	// })
	// router.Use(middlewareFunc)

	// Start serving the application
	logwrapper.Logger.Infoln("Running Server on port : ", serverConfig.Port)
	logwrapper.Logger.Fatal(router.Run(fmt.Sprintf(":%d", serverConfig.Port)))
}
