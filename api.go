package main

import (
	"encoding/hex"
	"time"

	"mtg/users"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type router struct{}

var Router router

const path = "/mtg"

type UserForm struct {
	Name   string `form:"name" json:"name"`
	Secret string `form:"secret" json:"secret"`
}

func InitGin() *gin.Engine {
	r := gin.New()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	users := r.Group(path)
	{
		users.GET("", Router.GetAll)
		users.POST("", Router.Create)
		users.DELETE(":name", Router.Delete)

	}
	return r
}

func (*router) Create(c *gin.Context) {
	var form UserForm

	if err := c.BindJSON(&form); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	u, err := users.User{
		Name: form.Name,
	}.Create()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "server error",
		})
		return
	}
	secret := "dd" + hex.EncodeToString(u.Secret)
	c.JSON(200, UserForm{Name: u.Name, Secret: secret})
}

func (*router) GetAll(c *gin.Context) {
	users, err := users.User{}.GetAll()
	var forms []UserForm

	if users == nil {
		c.JSON(404, gin.H{
			"error": "users not found",
		})
		return
	}

	if err != nil {
		c.JSON(404, gin.H{
			"error": "server error",
		})
		return
	}

	for _, u := range users {
		forms = append(forms, UserForm{u.Name, "dd" + hex.EncodeToString(u.Secret)})
	}

	c.JSON(200, forms)
}

func (*router) Delete(c *gin.Context) {
	name := c.Param("name")

	u := users.User{
		Name: name,
	}

	if !u.Exist() {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	err := u.Delete()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
