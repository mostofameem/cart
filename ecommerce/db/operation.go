package db

import (
	"ecommerce/model"
	"errors"
	"fmt"
)

func Register(name, email, pass string) error {
	dbpass := GetPass(name)
	if dbpass == "" {
		err := INSERT(name, email, pass)
		return err
	}
	return fmt.Errorf("user already exists")
}
func Login(email string, pass string) error {

	dbpass := GetPass(email)
	if dbpass == pass {
		return nil
	}
	return errors.New("failed ")
}
func AddToCart(item model.Cart, id string) error {

	product := GetProduct(item)
	query := "INSERT INTO cart values('" + id + "','" + product.Name + "','" + product.Price + "','" + item.Quantity + "');"
	_, err := Db.Exec(query)
	return err

}
func ShowCart(id string) ([]model.CartList, string) {
	listch := make(chan []model.CartList)
	totalch := make(chan string)

	go GiveMeTotal(id, totalch)
	go GiveMeCart(id, listch)

	list := <-listch
	total := <-totalch
	total = fmt.Sprintf("Total =%s", total)
	return list, total

}
