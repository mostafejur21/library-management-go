package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/mostafejur21/library-management-go/models"
	"github.com/mostafejur21/library-management-go/services"
)

var library services.LibraryService = &services.LibraryServices{}

func NewLibraryService() services.LibraryService {
	return &services.LibraryServices{
		Books:   []models.Book{},
		Members: []models.Member{},
	}
}

func AddBooks(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Book Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "available",
	}

	library.AddBooks(book)
	fmt.Print("Book added successfully!\n")
}

func RemoveBooks(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	bookIdStr, _ := reader.ReadString('\n')
	bookId, _ := strconv.Atoi(strings.TrimSpace(bookIdStr))

	library.RemoveBooks(bookId)
	fmt.Print("Book removed successfully!\n")
}

func ListBooks() {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}
