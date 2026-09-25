package main

import "fmt"

func main(){
	var x int
	var y int

	fmt.Printf("Ingrese el primer numero: ")
	fmt.Scan(&x)

	fmt.Printf("Ingrese el segundo numero: ")
	fmt.Scan(&y)

	if(x > y ){
		fmt.Println("El primer numero es mayor:",x, "es mayor que", y)
	} else if ( x < y) {
		fmt.Println("El segundo numero es mayor:",y, "es mayor que", x)
	} else if ( x == y ){
		fmt.Println("Son iguales")
	} else { 
		fmt.Println("Error!")
	}
}