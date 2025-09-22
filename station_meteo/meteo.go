package main

import (
	"fmt"
)

func MinimalTemperature(technicien_name string, temperature_matin int8, temperature_soir int8) {
	if temperature_matin < temperature_soir {
		fmt.Printf("la temperature minimal est: %d et la temperature maximale : %d, le techicien en charge est : %s\n", temperature_matin, temperature_soir, technicien_name)
	} else if temperature_matin > temperature_soir {
		fmt.Printf("la temperature minimal est: %d  et la temperature maximale : %d, le techicien en charge est : %s \n", temperature_soir, temperature_matin, technicien_name)
	}
}

func main() {
	var TM int8
	var TS int8
	var nom string
	fmt.Println("entre ton nom : ")
	fmt.Scanln(&nom)
	fmt.Println("entre ton année LA TEMPERATURE DU MATIN: ")
	fmt.Scanln(&TM)
	fmt.Println("entre ton année LA TEMPERATURE DU SOIR: ")
	fmt.Scanln(&TS)
	MinimalTemperature(nom, TM, TS)
}
