package routes

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type keyHolderV1 struct {
	Value int `bson:"key" json:"key"`
}

type keyHolderV2 struct {
	Value string `bson:"key" json:"key"`
}

// GetKeyV1 returns the encryption key
func (es *EchoServer) GetKeyV1(c echo.Context) error {

	if es.MessageKey.Value == 0 {
		es.MessageKey.Value = rand.Intn(100) + 1
	}

	return c.JSON(http.StatusOK,
		map[string]string{"key": strconv.Itoa(es.MessageKey.Value)})
}

// V2 key management system:

// // GetUserKey returns a key from the user.
// func (es *EchoServer) GetUserKey(c echo.Context) error {

// 	username := c.Param("name")

// 	userFromDB, err := es.Database.GetUser(username)
// 	if err != nil {
// 		es.Logger.Info("User not found in the database")

// 		return c.JSON(http.StatusNotFound,
// 			map[string]string{"result": "fail", "message": err.Error()})
// 	}

// 	es.Logger.Info("User found in the database")
// 	return c.JSON(http.StatusOK, userFromDB.Key)
// }

// // UpdateKey attempts to update an existing user's key in the database.
// func (es *EchoServer) UpdateKey(c echo.Context) error {

// 	username := c.Param("name")
// 	var tmpKey *keyHolder

// 	userFromDB, err := es.Database.GetUser(username)
// 	if err != nil {
// 		es.Logger.Info("User not found in the database")

// 		return c.JSON(http.StatusNotFound,
// 			map[string]string{"result": "fail", "message": err.Error()})
// 	}

// 	if err := c.Bind(&tmpKey); err != nil {
// 		es.Logger.Warn("Error binding incoming update key to keyHolder type", err)

// 		return c.JSON(http.StatusBadRequest,
// 			map[string]string{"result": "fail", "message": err.Error()})
// 	}

// 	if err := es.Database.UpdateUserKey(userFromDB.Name, tmpKey.Key); err != nil {
// 		es.Logger.Warn("Error writing new key to the database", err)

// 		return c.JSON(http.StatusInternalServerError,
// 			map[string]string{"result": "fail", "message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK,
// 		map[string]string{"result": "success", "message": "user key update success"})
// }
