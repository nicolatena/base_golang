package controllers


import (
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    . "go_rest_api/service/models"
)

func (idb *InDB) SelectDataUser(c *gin.Context) {

    var arr_data []User
    var arr_meta MetaUser
    var response ResponseUser

    idb.DB.AutoMigrate(&User{})

    idb.DB.Find(&arr_data)

    arr_meta = MetaUser{Status: true, Code: 200, Message: "Success"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}