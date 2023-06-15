package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var users = make(map[string]map[string]string)

const fileName = "users.txt"

func UserMangmentSystem() {
	fmt.Println("\nWelcome to Password Manager!")

	loadUsersFromFile()

	for {
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add User")
		fmt.Println("2. Show Usernames and Passwords")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			addUser()
		case 2:
			showUsernamesAndPasswords()
		case 3:
			saveUsersToFile()
			fmt.Println("3Thank you for using Password Manager!\n")
			return
		default:
			fmt.Println("\nInvalid choice. Please try again.")
		}
	}
}

func addUser() {
	fmt.Println("\nAdd User")
	fmt.Print("Enter username: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if username == "" {
		fmt.Println("Username cannot be empty.")
		return
	}

	if _, ok := users[username]; ok {
		fmt.Println("Username already exists.")
		return
	}

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" {
		fmt.Println("Password cannot be empty.")
		return
	}

	fmt.Print("Enter website or additional information: ")
	website, _ := reader.ReadString('\n')
	website = strings.TrimSpace(website)

	userData := map[string]string{
		"password": password,
		"website":  website,
	}

	users[username] = userData
	fmt.Println("User added successfully.")
}

func showUsernamesAndPasswords() {
	fmt.Println("\nUsernames and Passwords")
	for username, userData := range users {
		password := userData["password"]
		website := userData["website"]
		fmt.Printf("Username: %s, Password: %s, Website: %s\n", username, password, website)
	}
}

func loadUsersFromFile() {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 4 {
			username := parts[0]
			password := parts[1]
			website := parts[2]
			userData := map[string]string{
				"password": password,
				"website":  website,
			}
			users[username] = userData
		}
	}
}

func saveUsersToFile() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to save users to file.")
		return
	}
	defer file.Close()

	for username, userData := range users {
		password := userData["password"]
		website := userData["website"]
		line := fmt.Sprintf("%s:%s:%s\n", username, password, website)
		file.WriteString(line)
	}
}
