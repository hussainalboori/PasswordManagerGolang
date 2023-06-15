package main

import (
	"fmt"
)

func main() {
	for {
	var UserInput int
	fmt.Println("Please Choose From The Following\n1)Password Manager\n2)Password Generator\n3)Exit")
	fmt.Print("Your Input : ")
	fmt.Scanln(&UserInput)
	if UserInput == 1{
		UserMangmentSystem()
	}else if UserInput == 2{
		PasswordGenerator()
	}else if UserInput == 3 {
		break
	}else {
		fmt.Println("\nError The Input You Choose Not Available\n")
		continue
	}
}
}
