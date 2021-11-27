package main

import (
	"movie-api/driver"
	"movie-api/routes"
)

func main() {
	driver.InitDB()
	e := routes.New()

	// Starting the server
	e.Start(":8585")
}

// func UpdateAccount(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	reqAccount := Account{}
// 	c.Bind(&reqAccount)

// var dataAccount Account
// reqAccount.ID = uint(id)

// DB.Where("id = ?", id).First(&dataAccount)

// if reqAccount.Username != "" {
// 	dataAccount.Username = reqAccount.Username
// }

// if err := DB.Where("id = ?", id).Updates(&reqAccount).Error; err != nil {
// 	return c.JSON(http.StatusBadRequest, err)
// }
// SAVE FOR LATER
// var dataAccount Account
// var dataAccount ResponseAccount // COBA COBA
// DB.Raw("SELECT * FROM accounts WHERE id = ?", id).Find(&dataAccount)
// DB.Where("id = ?", id).Find(&dataAccount)
// reqAccount.ID = uint(id)

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "Success",
// 		"data":    dataAccount,
// 	})
// }

// func GetAccountByID(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	reqAccount := Account{}
// 	c.Bind(&reqAccount)

// 	if err := DB.Where("id = ?", id).First(&reqAccount); err != nil {
// 		return c.JSON(http.StatusNotFound, err)
// 	}

// 	var dataAccount ResponseAccount
// 	DB.Where("id = ?", id)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success",
// 		"data":    dataAccount,
// 	})
// if id <= len(data) && id > 0 {
// 	return c.JSON(http.StatusOK, data[id-1])
// } else {
// 	return c.JSON(http.StatusOK, data)
// }
//}
