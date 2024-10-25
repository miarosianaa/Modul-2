package main

import "fmt"

func main() {
        var bilangan int
        fmt.Print("Bilangan : ")
        fmt.Scan(&bilangan)

		if bilangan > 1 {
			fmt.Print("Faktor : ")
        	for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
					fmt.Print(i, " ")
                }
            }
        fmt.Println()
		} else {
			fmt.Println("Bilangan harus lebih dari 1")
		}
        
}