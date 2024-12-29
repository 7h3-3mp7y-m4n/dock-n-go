package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var user = []User{
	{"admin123", "$2a$10$wJPnALYw9t6yESNzU0vPlu0Zz7FysJLOxyI0V54lbWqY1X6GGoPOy"}, // password: password123 You should not hardcode user like this, use .enc to inject them
}

func login(ctx *gin.Context) {
	var loginUser User
	err := ctx.ShouldBindJSON(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	for _, value := range user {
		if value.UserName == loginUser.UserName {
			if err := bcrypt.CompareHashAndPassword([]byte(value.Password), []byte(loginUser.Password)); err == nil {
				token := genrateJWT(value.UserName)
				ctx.JSON(http.StatusAccepted, gin.H{"token": token})
				return
			}

		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
}

func genrateJWT(UserName string) string {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(cmd.jwtSecret)
	if err != nil {
		return ""
	}
	return tokenString
}
