package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-gin/models"
	"github.com/learn-gin/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user, Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse data.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message: ": "Invalid Credentials!", "error": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user.", "error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": jwtToken})
}
