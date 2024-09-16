package main

import "fmt"


func main () {
	var operasi_hitung string
	var angka_1, angka_2 float64
	
	fmt.Print("Input angka 1: ")
	fmt.Scanln(&angka_1)

	fmt.Print("Input angka 2: ")
	fmt.Scanln(&angka_2)

	fmt.Print("Input operasi hitung (+, -, *, /): ")
	fmt.Scanln(&operasi_hitung)


	switch operasi_hitung {
		case "+":
			fmt.Printf("%.3f %s %.3f = %.3f = %.3f", angka_1, operasi_hitung, angka_2, angka_1 + angka_2)
		case "-":
			fmt.Printf("%.3f %s %.3f = %.3f = %.3f", angka_1, operasi_hitung, angka_2, angka_1 - angka_2)
		case "*":
			fmt.Printf("%.3f %s %.3f = %.3f = %.3f", angka_1, operasi_hitung, angka_2, angka_1 * angka_2)
		case "/":
			fmt.Printf("%.3f %s %.3f = %.3f = %.3f", angka_1, operasi_hitung, angka_2, angka_1 / angka_2)
	
		default: 
		        fmt.Println("Kalkulator tidak mengerti input user")
	}

	fmt.Printf("%d %s %d = %d", angka_1, operasi_hitung, angka_2, output)


	





}
