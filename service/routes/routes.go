package routes

import (
    "github.com/gin-gonic/gin"
    "go_rest_api/service/config"
    "go_rest_api/service/controllers"
)

func Routes(route *gin.Engine){

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	users := route.Group("/v1/api/users")
	{
		users.GET("/select", inDB.SelectDataUser)
		// users.POST("/api/insert", inDB.InsertData)
		// users.PUT("/api/update/:id", inDB.UpdateData)
		// users.DELETE("/api/delete/:id", inDB.DeleteData)
	}
}