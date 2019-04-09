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

    idb.DB.Find(&arr_data)

    arr_meta = MetaUser{Status: true, Code: 200, Message: "Success"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) InsertDataUser(c *gin.Context) {
    
    var arr_data []User
    var arr_meta MetaUser
    var response ResponseUser
    sc_data := User{}


    err := c.BindJSON(&sc_data)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "can't bind struct",
        })
        return
    }

    sc_data.DeletedAt = nil
    
    idb.DB.Create(&sc_data)

    idb.DB.Find(&arr_data)

    arr_meta = MetaUser{Status: true, Code: 200, Message: "Success Inserted"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) UpdateDataUser(c *gin.Context) {
    
    var arr_data []User
    var arr_meta MetaUser
    var response ResponseUser
    sc_data := User{}

    
    id := c.Param("id")
    idb.DB.Where("id = ?", id).First(&sc_data)
    
    c.BindJSON(&sc_data)
    sc_data.DeletedAt = nil

    idb.DB.Save(sc_data)

    idb.DB.Find(&arr_data)

    arr_meta = MetaUser{Status: true, Code: 200, Message: "Success Updated"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}


func (idb *InDB) DeleteDataUser(c *gin.Context) {
    
    var arr_data []User
    var arr_meta MetaUser
    var response ResponseUser
    sc_data := User{}

    
    id := c.Param("id")
    idb.DB.Where("id = ?", id).First(&sc_data)
    idb.DB.Delete(sc_data)

    idb.DB.Find(&arr_data)

    arr_meta = MetaUser{Status: true, Code: 200, Message: "Success Deleted"}

    response.Meta = arr_meta
    response.Data = arr_data
 
    c.JSON(http.StatusOK, response)
}