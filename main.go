package main

import (
    "github.com/gin-gonic/gin"
    "rest-api-go/service/routes"
    
    "github.com/gin-contrib/static"
)


func main() {
	route := gin.Default()
	
	// inisialisasi folder upload
	route.Use(static.Serve("/service", static.LocalFile("rest-api-go/service", true)))

	routes.Routes(route)

	route.Run(":7777")
}
