package router

import (
	"net/http"

	mw "com.ashp8/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysessions", store))

	r.GET("/", home)

	r.POST("/login", handleLogin)
	r.POST("/signup", handleSignup)

	private := r.Group("/private")
	private.Use(mw.RequireAuth())
	{
		private.GET("/profile", getProfile)
		private.POST("/logout", handleLogOut)
	}

	return r
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
