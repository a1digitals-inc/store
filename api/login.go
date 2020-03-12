package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func Login(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "login",
		"bundle": "login",
	})
}

func PostLogin(c *gin.Context) {
	attempt := c.PostForm("password")
	password := db.GetPassword()
	// If there is no password in the database make the attempt the password
	if password == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(attempt), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
		}
		err = db.SetPassword(string(hash))
		if err != nil {
			log.Println(err)
		}
		c.JSON(200, gin.H{"data": "Password created"})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(attempt))
		if err != nil {
			c.JSON(200, gin.H{"valid": false})
			return
		}
		expiration := time.Now().Add(5 * time.Minute).Unix()
		claims := &jwt.StandardClaims{
			ExpiresAt: expiration,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET-KEY")))
		if err != nil {
			log.Println(err)
		}
		c.SetCookie("token", tokenString, 300, "/", os.Getenv("DOMAIN"), false, true)
		c.JSON(200, gin.H{"valid": true})
	}
}

func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		} else {
			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Parsing JWT token error")
				}
				return []byte(os.Getenv("SECRET-KEY")), nil
			})
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			if parsedToken.Valid {
				next(c)
			} else {
				c.JSON(401, gin.H{"error": "Invalid token"})
			}
		}
	}
}
