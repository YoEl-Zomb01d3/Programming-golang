package main

import "fmt"

var nombre string
var edad int
var altura float64
var peso float64
var imc float64

func main(){
	nombre = "Cristian"
	edad = 20
	altura = 1.70
	peso = 40

	imc = peso / (altura * altura)

	if(imc >= 30){
		fmt.Println("Estas gordo.")
	} else if(imc >= 25){
		fmt.Println("Estas mas o menos gordo.")
	} else if(imc >= 18){
		fmt.Println("Tas bien.")
	} else{
		fmt.Println("Estas bajo de peso, come mas.")
	}
}

