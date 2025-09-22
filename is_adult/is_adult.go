package main

import (
	"fmt"
)


func IsAdult(name string, naissance uint16, annee_actuel uint16){
	age := annee_actuel - naissance 
	if age < 18{
		fmt.Printf("Desolé, %s tu peux pas avoir de document officiel tu es encore mineur!\n", name)
	} else {
		fmt.Printf("Felicitation, %s tu peux avoir de document officiel tu es majeur! \n", name)
	}
}
func main(){
	AT := 2025
	var nom string 
	var naissance  uint16
	fmt.Println("entre ton nom : ")
	fmt.Scanln(&nom)
	fmt.Println("entre ton année de naissance: ")
	fmt.Scanln(&naissance)
	IsAdult(nom, naissance,uint16(AT))
}