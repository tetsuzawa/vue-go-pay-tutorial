package db

import (
	"fmt"
	"log"

	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/domain"
)

// SelectAllItems - select all posts
func SelectAllItems() (items domain.Items, err error) {
	log.Println("select all items called with request:", items)
	stmt, err := Conn.Query("SELECT * FROM items")
	if err != nil {
		return
	}
	defer stmt.Close()
	for stmt.Next() {
		var id int64
		var name string
		var description string
		var amount int64
		if err := stmt.Scan(&id, &name, &description, &amount); err != nil {
			continue
		}
		item := domain.Item{
			ID:          id,
			Name:        name,
			Description: description,
			Amount:      amount,
		}
		items = append(items, item)
	}
	return
}

// SelectItem - select post
func SelectItem(identifier int64) (item domain.Item, err error) {
	log.Println("select item called with request:", identifier)
	stmt, err := Conn.Prepare(fmt.Sprintf("SELECT * FROM items WHERE id = ? LIMIT 1"))
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	var id int64
	var name string
	var description string
	var amount int64
	err = stmt.QueryRow(identifier).Scan(&id, &name, &description, &amount)
	if err != nil {
		log.Println(err)
		return
	}
	item.ID = id
	item.Name = name
	item.Description = description
	item.Amount = amount
	return
}
