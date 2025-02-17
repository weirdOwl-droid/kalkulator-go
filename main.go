package main

import "fmt"

func main() {
	var (
		num1, num2 float64
		operator   string
	)

	fmt.Println("Silahkan masukkan angka ...")
	fmt.Scanln(&num1) // "&" digunakan sebagai untuk menyimpan input dari pengguna
	// fungsi (scan) digunakan untuk agar pengguna bisa menginputkan
	fmt.Println("Silahkan masukkan Operator : (+, -, *, /) ...")
	fmt.Scanln(&operator)
	fmt.Println("Silahkan memasukkan angka ...")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Hasilnya adalah\n %.0f", num1+num2)
	case "-":
		fmt.Printf("Hasilnya adalah\n %.0f", num1-num2)
	case "*":
		fmt.Printf("Hasilnya adalah\n %.0f", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Tidak bisa dibagi dengan 0")
		} else {
			fmt.Printf("Hasilnya adalah\n %.0f", num1/num2)
		}

	default:
		fmt.Println("Operator tidak ditemukan")
	}

}
