package main

import (
	"learn/httpserver/router"
	"learn/httpserver/setup"
	"os"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// at the begining init function get called
func init() {
	setup.LoadEnvVariable()
}

func main() {
	
	r := gin.Default()

	store := cookie.NewStore([]byte (os.Getenv("SESSION_KEY")))
	r.Use(sessions.Sessions("my-session",store))

	router.IndexRoute(r)
	
	r.Run() // Listen and server on port 8080

}
