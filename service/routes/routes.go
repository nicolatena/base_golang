package routes

import (
    "github.com/gin-gonic/gin"
    "rest-api-go/service/config"
    "rest-api-go/service/controllers"
)

func Routes(route *gin.Engine){

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		{
			users := api.Group("/users")
			{
				users.GET("/select", inDB.SelectDataUser)
				users.POST("/insert", inDB.InsertDataUser)
				users.PUT("/update/:id", inDB.UpdateDataUser)
				users.DELETE("/delete/:id", inDB.DeleteDataUser)
			}
		}
	}
}