package server

import (
	v1 "postinger/routes/v1"
)

// Initializing all the routes
func initializeRoutes() {

	// Version 1 routes
	V1Routes := router.Group("/api/v1/")
	{
		v1.ApiRoutes(V1Routes)
	}

}
