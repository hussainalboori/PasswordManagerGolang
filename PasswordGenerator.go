package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	numbers          = "0123456789"
	symbols          = "!@#$%^&*()-=_+[]{}|;':\"<>,.?/~`"
)

func generateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Create a combined string of characters to choose from
	allCharacters := uppercaseLetters + lowercaseLetters + numbers + symbols

	// Create a byte slice with the length of the desired password
	password := make([]byte, length)

	// Loop through each position in the password and randomly choose a character
	for i := 0; i < length; i++ {
		password[i] = allCharacters[rand.Intn(len(allCharacters))]
	}

	// Convert the byte slice to a string and return the password
	return string(password)
}

func PasswordGenerator() {
	var length int
	fmt.Print("Enter the length of the password: ")
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		fmt.Println("Invalid input. Exiting...")
		return
	}

	// Generate a random password with the user-specified length
	password := generateRandomPassword(length)
	fmt.Println("Random Password:", password)
	fmt.Println()
}
