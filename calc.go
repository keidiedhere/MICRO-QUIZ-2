package main

import "fmt"


func main () {
	var operasi_hitung string
	var angka_1, angka_2 int
	
	fmt.Print("Input angka 1: ")
	fmt.Scanln(&angka_1)

	fmt.Print("Input angka 2: ")
	fmt.Scanln(&angka_2)

	fmt.Print("Input operasi hitungan (+, -, *, /): ")
	fmt.Scanln(&operasi_hitung)

	output := 0

	switch operasi_hitung {
		case "+":
			output = angka_1 + angka_2
		case "-":
			output = angka_1 - angka_2
		case "*":
			output = angka_1 * angka_2
		case "/":
			output = angka_1 / angka_2
	
		default: 

		fmt.Println("kalkulator tidak mengerti input user")
	}

	fmt.Printf("%d %s %d = %d", angka_1, operasi_hitung, angka_2, output)


	





}
