package routes

import (
	"net/http"
	"time"
	
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
	"github.com/itsjamie/gin-cors"

    "rest-api-go/service/config"
	"rest-api-go/service/controllers"
)


func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        token := session.Get("token")
        if token == nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "GAGAL LOGIN",
            })
            c.Abort()
        } else {
            c.Next()
        }
    }
}


func Routes(route *gin.Engine){

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}


	route.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))


	store := sessions.NewCookieStore([]byte("secret"))
	route.Use(sessions.Sessions("mysession", store))

	v1 := route.Group("/v1")
	{

		
		auth := v1.Group("/auth")
		{
			auth.POST("/login", inDB.LoginHandler)
			auth.POST("/logout", inDB.LogoutHandler)
		}


		api := v1.Group("/api")
		{
			users := api.Group("/users")
			users.Use(AuthRequired())
			{
				users.GET("/select", inDB.SelectDataUser)
				users.POST("/insert", inDB.InsertDataUser)
				users.PUT("/update/:id", inDB.UpdateDataUser)
				users.DELETE("/delete/:id", inDB.DeleteDataUser)
			}
		}

	}
}