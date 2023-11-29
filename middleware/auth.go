package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser(c *gin.Context) {

	// token := c.Request.Header.Get("Authorization")
		
	//check session set or not 
		session := sessions.Default(c)
		fmt.Println(session.Get("isAuthenticated"))
		fmt.Println("loggedIn-token : ",session.Get("loggedInToken"))

		// fmt.Sprintf() convert interface type to string using %v - verb
		token := fmt.Sprintf("%v",session.Get("loggedInToken"))

	if token == "" {
		// c.String(http.StatusBadRequest, "Missing authorization header")
		c.String(http.StatusBadRequest, "Invalid request - please login to access")
		c.Abort() // prevent pending handler from being called
		return
	}

	err := verifyToken(token)

	if err != nil {
		c.String(http.StatusUnauthorized, "UnAuthorized")
		c.Abort()
		return
	}

	c.Next()
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		fmt.Errorf("Invalid token")
		return err
	}
	return nil
}

// ValidatePermission
func ValidatePermission(c *gin.Context) {

	role := c.Request.Header.Get("Role")

	fmt.Println(role)

	if role == "admin" {
		c.Next()
		return
	}

	c.String(http.StatusUnauthorized, "UnAuthorized")
	c.Abort()

}
