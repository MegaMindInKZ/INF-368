package a

import (
	"fmt"
	"strconv"
)

var Products []Product
var Users []User

type User struct {
	ID       int
	FullName string
}

type Product struct {
	ID     int
	Title  string
	UserID int
}

func AddUser(name string) {
	user := User{len(Users), name}
	Users = append(Users, user)
}

func AddProduct(title string, userID int) {
	product := Product{len(Products), title, userID}
	Products = append(Products, product)
}

func (user User) Print() {
	fmt.Println("ID: " + strconv.Itoa(user.ID) + ", Full Name: " + user.FullName)
}

func (product Product) Print() {
	fmt.Println("ID: " + strconv.Itoa(product.ID) + ", Title: " + product.Title)
}
