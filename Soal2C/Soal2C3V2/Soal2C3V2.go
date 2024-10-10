package main

import "fmt"

func main() {
	var bilangan, jumlahFaktor int

	fmt.Print("Bilangan : ")
	fmt.Scan(&bilangan)

	if bilangan > 0 {
		fmt.Print("Faktor : ")
		jumlahFaktor = 0
		for i := 1; i <= bilangan; i++ {
			if bilangan%i == 0 {
				fmt.Print(i, " ")
				jumlahFaktor++
			}
		}
		fmt.Println()

		if jumlahFaktor == 2 {
			fmt.Println("Prima : true")
		} else {
			fmt.Println("Prima : false")
		}
	} else {
		fmt.Println("Bilangan harus lebih dari 0")
	}
}
