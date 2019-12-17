package handler

import (
	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/db"
	"log"
	"net/http"
	"strconv"
)

// GetLists - get all items
func GetLists(c Context) error {
	res, err := db.SelectAllItems()
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, nil)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return err
	}
	err = c.JSON(http.StatusOK, res)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// GetItem - get item by id
func GetItem(c Context) error {
	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, nil)
		if err != nil {
			log.Fatalln(err)
			return err

		}
		return err
	}
	res, err := db.SelectItem(int64(identifier))
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, nil)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return err
	}
	err = c.JSON(http.StatusOK, res)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
