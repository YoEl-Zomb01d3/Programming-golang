package main

import "fmt"

func main(){
	var monto float64
	var descuento float64
	var total float64

	monto = 430

	if(monto >= 1000){
		descuento = monto * 0.20;
	} else if(monto >= 500){
		descuento = monto * 0.10;
	} else if(monto >= 100){
		descuento = monto * 0.05;
	} else{
		descuento = 0
	}

	total = monto - descuento

	fmt.Printf("Monto original: %.2f\n", monto)
	fmt.Printf("Descuento: %.2f\n", descuento)
	fmt.Printf("Total a pagar: %.2f\n", total)
}