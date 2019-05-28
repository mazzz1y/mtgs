package main
 
import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type router struct{}

var Router router

const path = "/mtg"

type UserForm struct {
	Name      string `form:"name" json:"name"`
	Secret    string `form:"secret" json:"secret"`
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
		// notes.PUT(":id", Router.Update)
		users.DELETE(":id", Router.Delete)

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

	u, err := User{
		Name: form.Name,
	}.Create()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "server error",
		})
		return
	}

	c.JSON(200, u)
}

func (*router) GetAll(c *gin.Context) {
	users, err := User{}.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "server error",
		})
		return
	}

	c.JSON(200, users)
}



func (*router) Delete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bad id",
		})
		return
	}

	u := User{
		ID: id,
	}

	if !u.Exist() {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	}

	if u.Delete(); err != nil {
		c.JSON(500, gin.H{
			"error": "server error",
		})
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}