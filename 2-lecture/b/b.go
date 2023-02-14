package main

import "github.com/MegaMindInKZ/a"

func main() {
	a.AddUser("Zanggar Zhumagaliyev")
	a.AddProduct("Jeans", 0)
	a.Users[0].Print()
	a.Products[0].Print()
}
