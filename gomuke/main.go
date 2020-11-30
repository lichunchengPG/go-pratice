package main

import (
	"fmt"
	"lichuncheng/gomuke/model"	
)

func main() {
	product := model.Product{}
	product.SetProductName("1231")
	fmt.Println(product.GetProductName())
}