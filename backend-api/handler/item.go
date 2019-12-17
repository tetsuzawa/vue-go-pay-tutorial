package handler

import (
	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/db"
	"log"
	"net/http"
)

// GetLists - get all items
func GetLists(c Context) {
	res, err := db.SelectAllItems()
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, nil)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	err = c.JSON(http.StatusOK, res)
	if err != nil {
		log.Fatalln(err)
	}
}

// GetItem - get item by id
func GetItem(c Context) {
	identifier := c.Param("id")
	res, err := db.SelectItem(identifier)
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, nil)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	err = c.JSON(http.StatusOK, res)
	if err != nil {
		log.Fatalln(err)
	}
}
