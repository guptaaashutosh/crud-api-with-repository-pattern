package controller

import (
	// "context"

	"fmt"
	"net/http"

	// "learn/httpserver/dal"
	// "learn/httpserver/model"
	"learn/httpserver/model"
	"learn/httpserver/repo"
	"learn/httpserver/setup"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	DB := setup.ConnectDB()
	//repositories initialization
	repos := repo.UserRepo(DB)
	getData, err := repos.GetData()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"AllData": getData,
	})
}

func Create(c *gin.Context) {
	DB := setup.ConnectDB()
	repos := repo.UserRepo(DB)
	//check data
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	isCreated, creationError := repos.CreateData(user)
	if creationError != nil {
		panic(creationError)
	}
	c.JSON(200, gin.H{
		"isCreated": isCreated,
	})

}

// Delete
func Delete(c *gin.Context) {
	DB := setup.ConnectDB()
	repos := repo.UserRepo(DB)
	id := c.Param("id")

	isCreated, deletionError := repos.DeleteData(id)
	if deletionError != nil {
		panic(deletionError)
	}
	c.JSON(200, gin.H{
		"isDeleted": isCreated,
	})

}

// Update
func Update(c *gin.Context) {
	DB := setup.ConnectDB()
	repos := repo.UserRepo(DB)

	id := c.Param("id")

	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	isUpdated, updationError := repos.UpdateData(user, id)
	if updationError != nil {
		panic(updationError)
	}
	c.JSON(200, gin.H{
		"isUpdated": isUpdated,
	})

}

// Login
func Login(c *gin.Context) {
	session := sessions.Default(c)
	//check if user has already loggedIn 
	if(session.Get("isAuthenticated") == true ) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "you are already loggedIn",
		})
		return
	}

	DB := setup.ConnectDB()
	repos := repo.UserRepo(DB)

	var loginUser model.Login
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	// check generate token after checking if it is available
	loggedInToken, loggedInError := repos.LoggedIn(loginUser)
	if loggedInError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": loggedInError.Error(),
		})
		return
	}

	//save authentication in session
	session.Set("isAuthenticated", true)
	session.Set("loggedInToken", loggedInToken)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"message": "successfully loggedIn",
	})

}

// Logout
func Logout(c *gin.Context) {
	//save authentication in session
	
	session := sessions.Default(c)

	if(session.Get("isAuthenticated") == false ){
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request to logout - already logout",
		})
		return 
	}

	session.Set("isAuthenticated", false)
	session.Set("loggedInToken", "")
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "successfully logout",
	})
}

// AuthData
func AuthData(c *gin.Context) {
	DB := setup.ConnectDB()
	//repositories initialization
	repos := repo.UserRepo(DB)
	getData, err := repos.GetData()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "successfully authenticated and authorized",
		"Auth Data": getData,
	})
}

// SessionTest
func SessionTest(c *gin.Context) {
	DB := setup.ConnectDB()
	//repositories initialization
	repos := repo.UserRepo(DB)
	getData, err := repos.GetData()
	if err != nil {
		panic(err)
	}

	// session management to check how many times the api get hit in current session
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	fmt.Println(v)
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"session-test-message": "successfully test session",
		"AllData":              getData,
	})
}
