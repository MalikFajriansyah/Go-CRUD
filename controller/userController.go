package controller

import (
	"go-crud/config"
	"go-crud/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all user
func GetUser(c echo.Context) error {
	// id := c.Param("ID")
	db := config.DB()

	var user []*model.Users
	if err := db.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	// if res := db.Find(&user, id); res.Error != nil {
	// 	data := map[string]interface{}{
	// 		"message": res.Error.Error(),
	// 	}
	// 	return c.JSON(http.StatusOK, data)
	// }

	// reponse := map[string]interface{}{
	// 	"data": user[0],
	// }

	return c.JSON(http.StatusOK, user)
}

//get by id

func GetById(c echo.Context) error {
	id := c.Param("id")
	var user model.Users
	db := config.DB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, user)
}

// add user
func AddUser(c echo.Context) error {
	var user []*model.Users
	db := config.DB()

	c.Bind(&user)
	db.Create(&user)

	return c.JSON(http.StatusOK, user)
	//Binding data
	// if err := c.Bind(u); err != nil {
	// 	data := map[string]interface{}{
	// 		"message": err.Error(),
	// 	}

	// 	return c.JSON(http.StatusInternalServerError, data)
	// }

	// user := &model.Users{
	// 	Fullname: u.Fullname,
	// 	Username: u.Username,
	// 	Email:    u.Email,
	// }

	// if err := db.Create(&user).Error; err != nil {
	// 	data := map[string]interface{}{
	// 		"message": err.Error(),
	// 	}
	// 	return c.JSON(http.StatusOK, data)
	// }

	// response := map[string]interface{}{
	// 	"data": u,
	// }
}

// update data user
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user model.Users
	db := config.DB()

	// if err := db.Where("id = ?", id).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// }
	// c.Bind(&user)
	// db.Save(&user)
	// return c.JSON(http.StatusOK, user)

	//binding data
	if err := c.Bind(user); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	exixting_user := new(model.Users)

	if err := db.First(&exixting_user, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	exixting_user.Fullname = user.Fullname
	exixting_user.Username = user.Username
	exixting_user.Email = user.Email

	if err := db.Save(&exixting_user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": exixting_user,
	}

	return c.JSON(http.StatusOK, response)
}

// delete user
func DeleteUser(c echo.Context) error {
	id := c.Param("ID")
	db := config.DB()

	user := new(model.Users)

	err := db.Delete(&user, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "user has deleted",
	}
	return c.JSON(http.StatusOK, response)
}
