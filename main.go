package main

import (
	"geoquest-backend/cmd/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
