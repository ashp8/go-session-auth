package router

import (
	"net/http"

	conn "com.ashp8/connection"
	"com.ashp8/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func handleLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := conn.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password!"})
		return
	}

	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login Successful!"})
}

func handleSignup(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err.Error()"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Email: input.Email, Password: string(hashedPassword)}

	if err := conn.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup Successful"})
}

func getProfile(c *gin.Context) {
	session := sessions.Default(c)
	var userId uint = session.Get("userId").(uint)

	var user models.User
	if err := conn.DB.First(&user, userId).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
	})

}

func handleLogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Logged Out"})
}
