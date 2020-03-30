package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"github.com/sergiosegrera/store/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func PostLogin(c *gin.Context) {
	var login models.Login
	c.Bind(&login)
	password := db.GetPassword()
	// If there is no password in the database make the attempt the password
	if password == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.MinCost)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		err = db.SetPassword(string(hash))
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Password created"})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
		if err != nil {
			c.JSON(401, gin.H{"message": "Wrong password"})
			return
		}
		expiration := time.Now().Add(30 * time.Minute).Unix()
		claims := &jwt.StandardClaims{
			ExpiresAt: expiration,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		c.SetCookie("token", tokenString, 1800, "/", os.Getenv("DOMAIN"), false, false)
		c.JSON(200, gin.H{"message": "Authenticated"})
	}
}

func PostRefresh(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not refresh token"})
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Parsing JWT token error")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(500, gin.H{"message": "Error decoding token"})
		return
	}
	if parsedToken.Valid {
		expiration := time.Now().Add(30 * time.Minute).Unix()
		claims := &jwt.StandardClaims{
			ExpiresAt: expiration,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		c.SetCookie("token", tokenString, 1800, "/", os.Getenv("DOMAIN"), false, false)
		c.JSON(200, gin.H{"message": "Refreshed"})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not refresh token"})
		return
	}
}

func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(302, "/login")
			return
		} else {
			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Parsing JWT token error")
				}
				return []byte(os.Getenv("SECRET_KEY")), nil
			})
			if err != nil {
				c.JSON(500, gin.H{"message": err.Error()})
				return
			}
			if parsedToken.Valid {
				next(c)
			} else {
				c.JSON(301, gin.H{"message": "Invalid Token"})
			}
		}
	}
}
