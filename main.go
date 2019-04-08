package main

import (
    "github.com/gin-gonic/gin"
    "go_rest_api/service/routes"
    
    "github.com/gin-contrib/static"
    "github.com/foolin/gin-template"
)


func main() {
	route := gin.Default()
	
	// inisialisasi folder upload
	route.Use(static.Serve("/service", static.LocalFile("go_rest_api/service", true)))

	routes.Routes(route)

	route.Run(":7777")
}
