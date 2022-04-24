package controllers

/**
Controller holds the business logic to perform after the request was received,
and it talks to the model layer to make any DB changes that may be required.
**/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhisenberg/bookstore/pkg/models"
	"github.com/abhisenberg/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

/**
Function to return all books in the DB
Params: w ReponseWriter, r HttpRequest
**/

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)                   //Convert the model into JSON (known as marshalling)
	w.Header().Set("Content-Type", "pkglication/json") //❗❗ What is pkglication? Why is it used?
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       //Getting all the vars (path parameters) from the URL using mux.vars
	bookId := vars["bookId"]                  //Taking the bookId param from the set of vars
	Id, err := strconv.ParseInt(bookId, 0, 0) //Usually the params are of type str, convert it to int to make the query
	if err != nil {
		fmt.Printf("Error while parsing: %v\n", err)
	}
	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}   // ❗❗ Why use & here? What happens if we don't?
	utils.ParseBody(r, createBook) //Parses the request body to get the book content, unmarashall it into the book object "createBook"
	b := createBook.CreateBook()   //Adds the new book object in the DB
	res, _ := json.Marshal(b)      //Now we will return the newly added object to the JSON, hence we marshal it
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0) //Getting the bookId to delete from the request, convert it into int
	if err != nil {
		fmt.Printf("Error while parsing: %v\n", err)
	}
	book := models.DeleteBook(id) //Delete the book from the DB, it returns the deleted book
	res, _ := json.Marshal(book)  //Prepare to send the deleted-book info back to the user, hence marshal it
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} //Create empty book to store the updated-book information
	utils.ParseBody(r, updateBook)  //Parse the request to get the content of the updated-book

	vars := mux.Vars(r) //Now, getting the ID of the book that needs to be deleted
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing: %v\n", err)
	}

	//Find the book to delete in the DB
	bookDetails, db := models.GetBookById(id)

	//Update the model with the new values if present
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	//Save it in the DB
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
