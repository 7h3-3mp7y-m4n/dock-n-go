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

var users = []User{}

func Login(ctx *gin.Context) {
	var loginUser User
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	for _, value := range users {
		if value.UserName == loginUser.UserName {
			if err := bcrypt.CompareHashAndPassword([]byte(value.Password), []byte(loginUser.Password)); err == nil {
				token := generateJWT(value.UserName)
				ctx.JSON(http.StatusAccepted, gin.H{"token": token})
				return
			}
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
}
func Register(ctx *gin.Context) {
	var newUser User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)
	users = append(users, newUser)

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func generateJWT(UserName string) string {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("Dock-n-go"))
	if err != nil {
		return ""
	}
	return tokenString
}
