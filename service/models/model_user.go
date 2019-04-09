package models


import (
  "github.com/jinzhu/gorm"
)



type User struct {
  gorm.Model
  Email string `json:"email" form:"email" gorm:"unique"`
  Password string `json:"password" form:"password"`
  Role string `json:"role"`
  Isactivated int `json:"isactivated" form:"isactivated"`
  Isused int `json:"isused" form:"isused"`
}


type MetaUser struct {
    Status bool `json:"status"`
    Code int `json:"code"`
    Message string `json:"message"`
}

type ResponseUser struct {
    Meta MetaUser `json:"meta"`
    Data []User `json:"data"`
}