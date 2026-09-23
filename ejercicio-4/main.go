package main

import "fmt"

func main(){
	var x int

	fmt.Printf("Ingrese el numero: ");
	fmt.Scan(&x)
	
	if(x % 2 == 0){
		fmt.Println("El numero es par")
	} else{
		fmt.Println("El numero es impar")
	}

	
}