package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json: "id"`
	Email    string `json: "email"`
	Nama     string `json: "name"`
	UserName string `json: "username"`
}

// inisiasi db
var users []User

func main() {
	e := echo.New()

	//routing
	e.GET("/users", getUser)
	e.GET("/users/:id", getUserById)
	e.POST("/users/create", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Start(":8080")
}

// get all user
func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// get user by id
func getUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, "User Not Found")
}

// create data user
func createUser(c echo.Context) error {
	//inisiasi data baru
	user := new(User)

	if err := c.Bind(user); err != nil {
		return err
	}

	//automatically id
	user.Id = len(users) + 1
	users = append(users, *user)

	return c.JSON(http.StatusCreated, user)
}

// update data user
func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range users {
		if users[i].Id == id {
			updatedUser := new(User)
			if err := c.Bind(updatedUser); err != nil {
				return err
			}

			// Update student data
			users[i].Nama = updatedUser.Nama
			users[i].Email = updatedUser.Email
			users[i].UserName = updatedUser.UserName

			return c.JSON(http.StatusOK, users[i])
		}
	}
	return c.JSON(http.StatusNotFound, "User Not Found")
}

// delete data user
func deleteUser(c echo.Context) error {
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range users {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "User Not Found")
}
