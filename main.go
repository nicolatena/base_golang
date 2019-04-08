package main

import (
    "github.com/gin-gonic/gin"
    "go_rest_api/service/routes"
    
    "github.com/gin-contrib/static"
    "github.com/foolin/gin-template"
)


func main() {
	route := gin.Default()

	// untuk mendeteksi javascript, css, image di puclic
	route.Use(static.Serve("/frontend", static.LocalFile("go_rest_api/frontend", true)))
	// inisialisasi folder upload
	route.Use(static.Serve("/service", static.LocalFile("go_rest_api/service", true)))
	// untuk render file template HTML
	route.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "go_rest_api/frontend/view/html",
		Extension:    ".tmpl",
		Master:       "support/master",
		DisableCache: true,
	})

	routes.Routes(route)

	route.Run(":7777")
}
