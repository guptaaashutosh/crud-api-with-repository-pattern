package repo

import (
	"context"
	"fmt"
	"learn/httpserver/model"
	"learn/httpserver/utils"
	"log"
	"strconv"
)

type Repositories interface {
	GetData() ([]model.User, error)
	CreateData(model.User) (bool, error)
	DeleteData(string) (bool, error)
	UpdateData(model.User, string) (bool, error)
	LoggedIn(model.Login) (string, error)
}

func (u User) GetData() ([]model.User, error) {
	//call dal to get data
	rows, err := u.db.Query(context.Background(), "SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}

	var allData []model.User

	for rows.Next() {
		var data model.User
		err := rows.Scan(&data.Id, &data.Name, &data.Age, &data.Address)
		if err != nil {
			fmt.Println(err)
		}
		allData = append(allData, data)
		fmt.Println(data)
	}

	defer rows.Close()
	return allData, nil

}

func (u User) CreateData(data model.User) (bool, error) {
	var isCreated bool
	//call dal to get data
	fmt.Println(data)
	createdData, err := u.db.Query(context.Background(), "INSERT INTO employee(id,name,age,address)values($1,$2,$3,$4) RETURNING *", data.Id, data.Name, data.Age, data.Address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createdData)
	if createdData != nil {
		isCreated = true
	} else {
		isCreated = false
	}

	return isCreated, nil

}

// DeleteData
func (u User) DeleteData(id string) (bool, error) {
	var isDeleted bool
	deletedData, err := u.db.Query(context.Background(), "DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deletedData)
	if deletedData != nil {
		isDeleted = true
	} else {
		isDeleted = false
	}

	return isDeleted, nil

}

// UpdateData
func (u User) UpdateData(updateData model.User, id string) (bool, error) {
	var isUpdated bool
	updatedData, err := u.db.Query(context.Background(), "UPDATE employee SET name=$2,age=$3,address=$4 WHERE id=$1", id, updateData.Name, updateData.Age, updateData.Address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedData)
	if updatedData != nil {
		isUpdated = true
	} else {
		isUpdated = false
	}

	return isUpdated, nil

}

// LoggedIn
// func (u User) LoggedIn(user model.Login) (bool, error) {
// 	var isLoggedIn bool
// 	loggedIn, err := u.db.Query(context.Background(), "SELECT * FROM employee WHERE id=$1 AND name=$2", user.Id, user.Name)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	fmt.Println(loggedIn.CommandTag().RowsAffected())

// 	if loggedIn != nil {
// 		isLoggedIn = true
// 	} else {
// 		isLoggedIn = false
// 	}

// 	return isLoggedIn, nil

// }

func (u User) LoggedIn(user model.Login) (string, error) {

	_, err := u.db.Query(context.Background(), "SELECT * FROM employee WHERE id=$1 AND name=$2", user.Id, user.Name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//generate token
	loggedInToken, tokenError := utils.GenerateToken(strconv.FormatInt(user.Id,10))
	if tokenError != nil {
		return "",tokenError
	}

	return loggedInToken, nil

}
