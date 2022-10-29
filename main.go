package main

import (
	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/routes"
	"github.com/proyecto-final-2022/geoquest-backend/config"

	_ "github.com/proyecto-final-2022/geoquest-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title GeoQuest Backend
//@version 1.0
//@description App designed to communicate with GeoQuest's mobile app and provide CRUD functions
func main() {

	r := gin.Default()
	config.Connect()

	router := routes.NewRouter(r)
	router.MapRoutes()

	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

/*
	done := make(chan bool)
	go forever()
	<-done // Block forever
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
*/
