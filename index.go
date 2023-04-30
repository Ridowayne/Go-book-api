package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
 type book struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int64 `json:"quantity"`
 }

 var books = []book{
	{ID: 1, Title: "Richest man in Babylon", Author: "George Samuel clason", Quantity:8},
	{ID: 2, Title: "Rich Dad, Poor Dad", Author: "Robert Kiyosaki", Quantity:9},
	{ID: 3, Title: "Dreams from my father", Author: "Barack Obama", Quantity:9},
	{ID: 4, Title: "Half of a yellow Sun", Author: "Chimamanda Ngozi Adichie", Quantity:25},
	{ID: 5, Title: "The secrete lives Baba Segi's  Wives", Author: "Lola shoneyin", Quantity:14},
	{ID: 6, Title: "Purple Hibiscus", Author: "Chimamanda Ngozi Adichie", Quantity:99},
 }

 func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

 }
  func createBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, books)
  }
  func bookById(c *gin.Context) {
	id:=c.Param("id")
	book, err:= getBookById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
	}
	c.IndentedJSON(http.StatusOK, book)
  }
  func getBookById(id string) (*book, error) {
	
	for i, b := range books{
		iDD, _:= strconv.ParseInt(id, 10, 64)
		if b.ID == iDD {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")

  }
  func checkOutBook(c *gin.Context) {
	id, ok:= c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id"}) 
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error getting book"})
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
  }
  func returnBook (c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
	}
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
  }

func main () {
	fmt.Println("Application starting...")
	router:= gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books", createBooks)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkOut", checkOutBook)
	router.PATCH("/checkIn", returnBook)
	router.Run("localhost:8080")

}