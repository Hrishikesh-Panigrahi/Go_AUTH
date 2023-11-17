package middleware

import (
	"GO_AUTH/initializers"
	"GO_AUTH/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])	

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			
		}

		c.Set("user", user)
		c.Next()
		// fmt.Println(, claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}
