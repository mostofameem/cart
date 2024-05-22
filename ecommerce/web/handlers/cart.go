package handlers

import (
	"ecommerce/auth"
	"ecommerce/db"
	"ecommerce/model"
	"ecommerce/web/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	item, err := UrlOperation(r.URL.String())
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Error Loading Query")
		return
	}

	userID, err := GetIdFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		utils.SendError(w, 404, err, "Invalid Id")
	}

	err = db.AddToCart(item, userID)
	if err != nil {
		utils.SendError(w, 404, err, "Add to cart Failed")
		return
	}
	utils.SendData(w, "Added To cart successful")
}
func UrlOperation(r string) (model.Cart, error) {

	var item model.Cart
	parsedUrl, err := url.Parse(r)
	if err != nil {
		return item, err
	}
	queryParams := parsedUrl.Query()
	item.ProductName = queryParams.Get("product_name")
	item.Quantity = queryParams.Get("quantity")

	return item, nil
}
func GetIdFromHeader(r string) (string, error) {

	authHeader := r
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := auth.ParseToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("error Parsing")
	}

	// Extract user ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, _ := claims["Id"].(string)
	return userID, nil
}

func ShowCart(w http.ResponseWriter, r *http.Request) {

	id, err := GetIdFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		utils.SendError(w, 404, err, "Payload Error")
	}
	list, total := db.ShowCart(id)
	utils.SendBothData(w, total, list)
}
