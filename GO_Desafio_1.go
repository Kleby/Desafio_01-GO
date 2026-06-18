package main

import "fmt"

func conversorTerm(grausKelvin float64) float64 {
	return grausKelvin - 273.0
}

func inputF64() float64 {
	var number float64 = 0
	fmt.Scan(&number)
	return number
}

func main() {
	fmt.Println("Informe a temperatura em graus Kelvin: ")
	k := inputF64()

	fmt.Printf("A temperatura em %gº é %gk ", k, conversorTerm(k))
}
