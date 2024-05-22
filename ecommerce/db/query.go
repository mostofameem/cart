package db

import "ecommerce/model"

func GetPass(email string) string {
	query := "SELECT PASSWORD from users where email ='" + email + "';"

	var password string
	Db.QueryRow(query).Scan(&password)

	return password

}
func INSERT(name, email, pass string) error {

	query := "INSERT into users (name,email,password) VALUES ('" + name + "','" + email + "', '" + pass + "');"
	_, err := Db.Exec(query)
	return err
}
func GetUser(email string, usrchan chan model.User) {

	query := "SELECT id, name, email FROM users WHERE email = '" + email + "';"
	var user model.User

	err = Db.QueryRow(query).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		usrchan <- model.User{}
		close(usrchan)
		return
	}

	usrchan <- user
	close(usrchan)
}
func GetProduct(item model.Cart) model.Product {
	query := "SELECT name,price,quantity from products where name='" + item.ProductName + "';"
	var product model.Product
	err = Db.QueryRow(query).Scan(&product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return model.Product{}
	}
	return product
}
func GiveMeCart(id string, ch chan []model.CartList) {
	query := "SELECT product_name,quantity,price from cart where user_id=" + id + ";"
	rows, err := Db.Query(query)
	if err != nil {
		ch <- []model.CartList{}
	}

	var AllProduct []model.CartList

	for rows.Next() {
		var Product model.CartList
		err := rows.Scan(&Product.ProductName, &Product.Quantity, &Product.Price)
		if err != nil {
			ch <- []model.CartList{}
		}
		AllProduct = append(AllProduct, Product)
	}
	ch <- AllProduct
}
func GiveMeTotal(id string, totalchan chan string) {
	query := "SELECT sum(price*quantity) from cart where user_id=" + id + ";"
	var total string
	Db.QueryRow(query).Scan(&total)

	totalchan <- total
}
