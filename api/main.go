package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/item", getUsers)

	e.POST("/item", createUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
	return json.NewEncoder(c.Response().Writer).Encode(items)
}

func createUser(c echo.Context) error {
	var item Item

	if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	items = append(items, item)

	return c.String(http.StatusCreated, "User  created successfully")
}

type Item struct {
	Caption string  `json:"caption"`
	Weight  float32 `json:"weight"`
	Number  int     `json:"number"`
}

var items = []Item{}

// items = []Item{
//     {Caption: "Item 1", Weight: 1.0, Number: 1},
//     {Caption: "Item 2", Weight: 2.0, Number: 2},
//     {Caption: "Item 3", Weight: 3.0, Number: 3},
// }
