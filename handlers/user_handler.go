package handlers

import (
	"net/http"
	"time"

	"github.com/Dnreikronos/budgetMannager---Back/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("JWT_SECRET")

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(user models.User) (string, error) {
  expirationTime := time.Now().Add(24 * time.Hour)
  claims := &jwt.StandardClaims{
    Subject: user.ID.String(),
    ExpiresAt: expirationTime.Unix(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(jwtSecret)
}

func CreateUserHandler(c *gin.Context) {
  var input models.SignInInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  hashedPassword, err := HashPassword(input.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
    return
  }
  newUser := models.User{
    Email: input.Email,
    Password: hashedPassword,
    Verified: true,
  }
  db := c.MustGet("db").(*gorm.DB)
  if err := db.Create(&newUser).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
  }
  c.JSON(http.StatusCreated, models.FilteredResponse(newUser))
}

func LoginHandler(c *gin.Context) {
  var input models.SignInInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  var existingUser models.User
  db := c.MustGet("db").(*gorm.DB)
  if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }
  if err := VerifyPassword(existingUser.Password, input.Password); err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  } 
  token, err := GenerateToken(existingUser)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"token": token})
}


