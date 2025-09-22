package main 

import ( 
	"fmt"
)

func Invoice(product_price ,client_money  float32){
	if product_price < client_money {
	fmt.Printf("ton argent depasse le prix du produit, ton relicat : %f \n", client_money - product_price)
	}else if  product_price > client_money {
	fmt.Printf("ton argent est inferieur au prix du produit, tu dois completer : %f \n", product_price - client_money)
	}else if  product_price == client_money {
	fmt.Printf("parfait tu as remis exactement le prix dui produits \n")
	}
}

func main() {
	var CM float32
	var price float32
	fmt.Print("tu as gardé combien :")
	fmt.Scanln(&CM)
	price =54.50

	Invoice(price, CM)
}