package main

import "fmt"

func main(){
	var letra string

	fmt.Printf("Ingrese la letra: ")
	fmt.Scan(&letra)

	if (letra == "a" || letra == "e" || letra == "i" || letra == "o" || letra == "u") {
    	fmt.Println("Es vocal")
	} else {
    	fmt.Println("No es vocal")
	}
}