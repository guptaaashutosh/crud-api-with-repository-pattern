package controller

import (
	// "context"

	"net/http"

	// "learn/httpserver/dal"
	// "learn/httpserver/model"
	"learn/httpserver/model"
	"learn/httpserver/repo"
	"learn/httpserver/setup"

	"github.com/gin-gonic/gin"
)

// func Test(c *gin.Context) {
// 	data := map[string]interface{}{
// 		"id":   10,
// 		"name": "test-name",
// 		"age":  30,
// 	}
// 	c.JSON(http.StatusOK, data)

// 	// c.JSON(http.StatusOK, gin.H{  // gin.H helps to access the element. gin.H is defined as type H map[string]interface{}.
// 	// 	"message": "pong",
// 	//   })
// }

// // func (r *repo.User) Get() model.User{
// //     return dal.Get(r)
// // }

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

//Delete
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

//Update
func Update(c *gin.Context) {
	DB := setup.ConnectDB()
	repos := repo.UserRepo(DB)

	id := c.Param("id")

	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	isUpdated, updationError := repos.UpdateData(user,id)
	if updationError != nil {
		panic(updationError)
	}
	c.JSON(200, gin.H{
		"isUpdated": isUpdated,
	})

}

