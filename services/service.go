package services

import (
	"fmt"

	"github.com/mostafejur21/library-management-go/models"
)

type LibraryService interface {
	AddBooks(book models.Book)
	RemoveBooks(bookID int)
	// BorrowedBooks(bookID, memberID int) error
	// ReturnBooks(bookID, memberID int) error
	ListAvailableBooks() []models.Book
	// ListBorrowedBooks() []models.Book
}

type LibraryServices struct {
	Books   []models.Book
	Members []models.Member
}

func (service *LibraryServices) AddBooks(book models.Book) {
	service.Books = append(service.Books, book)
}

func (service *LibraryServices) RemoveBooks(bookID int) {
	for i, book := range service.Books {
		if book.ID == bookID {
			service.Books = append(service.Books[:i], service.Books[i+1:]...)
			return
		}
	}
}

func (service *LibraryServices) ListAvailableBooks() []models.Book {
	var availableBook []models.Book = make([]models.Book, 0)
	for _, book := range service.Books {
		fmt.Println(book)
		if book.Status == "available" {
			availableBook = append(availableBook, book)
		}
	}
	return availableBook

}
