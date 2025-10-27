package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mostafejur21/library-management-go/controllers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add new books")
		fmt.Println("2. Remove books")
		fmt.Println("3. List all books")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			controllers.AddBooks(reader)
		case "2":
			controllers.RemoveBooks(reader)
		case "3":
			controllers.ListBooks()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
