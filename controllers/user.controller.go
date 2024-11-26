package controllers

import (
	"fmt"
	"net/http"

	"github.com/bebek-goreng/golang-jwt-auth/initializer"
	"github.com/bebek-goreng/golang-jwt-auth/models"
	"github.com/bebek-goreng/golang-jwt-auth/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if user.FirstName == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input - FirstName, Email, & Password required"})
		return
	}

	var existingUser models.User

	if err := initializer.DB.First(&existingUser, "email = ?", user.Email).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already use"})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Failed to hash password"})
	}

	user.Password = string(hashedPassword)

	if err := initializer.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func SignIn(c *gin.Context) {
	var userInput struct {
		Email    string
		Password string
	}

	var user models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input - Email and password required"})
		return
	}

	initializer.DB = initializer.DB.Debug()

	if err := initializer.DB.First(&user, "email = ?", userInput.Email).Error; err != nil {
		fmt.Println("Error:", err)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		}

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, err := utils.GenerateToken(uint(user.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Failed to generate token"})
		return
	}

	c.SetCookie("access_token", accessToken, 3600*24, "/", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})

}
