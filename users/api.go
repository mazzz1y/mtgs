package users

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"mtgs/config"

	"go.uber.org/zap"
)

// Creates type for API methods
type apiMethods struct{}

var (
	api    apiMethods
	log, _ = zap.NewProduction()
)

// UserForm defines JSON form for API
type UserForm struct {
	Name   string `form:"name" json:"name"`
	Secret string `form:"secret" json:"secret"`
}

// StartAPI creates a new Gin instance and start it in a goroutine
func StartAPI(conf *config.Config) {
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

	if conf.APIToken != "" {
		r.Use(authMiddleware(conf.APIToken))
	}

	users := r.Group(conf.APIBasepath)
	{
		users.GET("", api.getAll)
		users.POST("", api.create)
		users.DELETE(":name", api.delete)
	}

	apiListenAddr := conf.BindIP.String() + ":" + strconv.Itoa(int(conf.BindAPIPort))
	InitKV(conf)
	go r.Run(apiListenAddr)
}

func (*apiMethods) create(c *gin.Context) {
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
		zap.S().Errorw("Failed to create user", "error", err)
		return
	}
	secret := "dd" + hex.EncodeToString(u.Secret)
	c.JSON(200, UserForm{Name: u.Name, Secret: secret})
}

func (*apiMethods) getAll(c *gin.Context) {
	users, err := User{}.GetAll()
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
		zap.S().Errorw("Failed to get users", "error", err)
		return
	}

	for _, u := range users {
		forms = append(forms, UserForm{u.Name, "dd" + hex.EncodeToString(u.Secret)})
	}

	c.JSON(200, forms)
}

func (*apiMethods) delete(c *gin.Context) {
	name := c.Param("name")

	u := User{
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
		zap.S().Errorw("Failed to delete users", "error", err)
		return
	}

	c.JSON(200, gin.H{
		"status": "deleted",
	})
}

func authMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != token {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}
	}
}
