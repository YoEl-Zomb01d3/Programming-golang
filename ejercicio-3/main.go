package main

import "fmt"

func main(){
	var temperatura float64

	fmt.Print("Ingrese la temperatura en °C: ")
	fmt.Scan(&temperatura)

	if(temperatura >= 40){
		fmt.Println("Extremo")
	} else if(temperatura >= 30){
		fmt.Println("Caluroso")
	} else if(temperatura >= 20){
		fmt.Println("Templado")
	} else if(temperatura >= 10){
		fmt.Println("Frio")
	} else if(temperatura >= 0){
		fmt.Println("Muy frio")
	} else{
		fmt.Println("Congelante")
	}
}