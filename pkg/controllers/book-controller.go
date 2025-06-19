package controllers

import (
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/abh1SHAKE/go-crud/pkg/utils"
	"github.com/abh1SHAKE/go-crud/pkg/models"
	"github.com/abh1SHAKE/go-crud/pkg/repository"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := repository.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	book, err := repository.GetBookById(bookId)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

 	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	err := utils.ParseBody(r, book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = repository.CreateBook(book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	book, err := repository.DeleteBook(bookId)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) 
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	err := utils.ParseBody(r, updateBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Error while parsing", http.StatusBadRequest)
		return
	}

	book, err := repository.GetBookById(bookId)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	err = repository.UpdateBook(book)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}