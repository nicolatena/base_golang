package controllers


import (
    "log"
    "net/http"
    "strings"
    "crypto/md5"
    "encoding/hex"

    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/gin-gonic/contrib/sessions"
    jwt "github.com/dgrijalva/jwt-go"
    . "rest-api-go/service/models"
)


func (idb *InDB) LoginHandler(c *gin.Context) {
    var user User
    var arr_user User
    sc_data := User{}


    sc_data.DeletedAt = nil
    sc_data.Email = "admin"
    sc_data.Password = "21232f297a57a5a743894a0e4a801fc3"
    idb.DB.Create(&sc_data)
    
    session := sessions.Default(c)

    err := c.Bind(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "can't bind struct",
        })
    }
    hasher := md5.New()
    hasher.Write([]byte(user.Password))
    password := hex.EncodeToString(hasher.Sum(nil))

    query := idb.DB.Where("LOWER(email) = ? AND password = ?", strings.ToLower(user.Email), password).First(&arr_user)

    if query.RowsAffected <= 0 {
        c.JSON(http.StatusUnauthorized,  gin.H{
            "status":  http.StatusUnauthorized,
            "message": "wrong username or password",
        })
        return
    }else{
        sign := jwt.New(jwt.GetSigningMethod("HS512"))

        claims := sign.Claims.(jwt.MapClaims)
        claims["foo"] = "bar"
        claims["user"] = user.Email
        
        token, err := sign.SignedString([]byte("secret"))

        session.Set("token", token) //In real world usage you'd set this to the users ID
        err = session.Save()

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "message": err.Error(),
            })
            c.Abort()
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Login Success",
            "token": token,
        })
    }
}

func (idb *InDB) LogoutHandler(c *gin.Context) {

    session := sessions.Default(c)
    token := session.Get("token")
    if token == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
    } else {
        log.Println(token)
        session.Delete("token")
        session.Save()
        c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
    }
}