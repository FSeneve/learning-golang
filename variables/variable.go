package main

import "fmt"

func main() {
	var age int = 30
	var prenom = "Francis"

	isActive := true

	fmt.Println(age)
	fmt.Println(prenom)

	if age > 18 {
		fmt.Println("Adulte")
	} else {
		fmt.Println("Tu es mineur(e)")
	}

	if isActive {
		fmt.Println("Connected")
	} else {
		fmt.Println("Check your credentials")
	}
}
