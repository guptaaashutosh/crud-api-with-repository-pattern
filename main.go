package main

import (
	// "context"
	// "fmt"

	"github.com/gin-gonic/gin"
	"learn/httpserver/router"
	"learn/httpserver/setup"
	// "log"
	// "os"
)

// at the begining init function get called
func init() {
	setup.LoadEnvVariable()
	//database connection

}

func main() {

	
	r := gin.Default()
	router.IndexRoute(r)
	
	r.Run() // Listen and server on port 8080

	
	// })
	// r := gin.Default() // here r is a router
	// r.GET("/test", func(c *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"id":   10,
	// 		"name": "test-name",
	// 		"age":  30,
	// 	}
	// 	c.JSON(200, data)

	//repository pattern
	// DB := setup.ConnectDB()

	

	//send repos to controller 

	
	// DB := setup.ConnectDB()



	// fetchData(conn)

	// rows, err := DB.Query(context.Background(), "SELECT * FROM employee")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // fmt.Println(rows.scan())
	// for rows.Next() {
	// 	var name string
	// 	var address string
	// 	var id, age int
	// 	err = rows.Scan(&id, &name, &age, &address)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	fmt.Println(id, name, age, address)
	// }

	// defer rows.Close()
	// // defer DB.Close(context.Background())

	// router.Run(":9000")   // Listen and server on port 9000

}
