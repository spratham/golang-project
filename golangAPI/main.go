package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model
type Books struct {
	BookID    string `json:"BookID"`
	Book      string `json:"book"`
	BookPrice int    `json:"price"`
}

var books []Books

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//data input
	books = append(books, Books{BookID: "1", Book: "Galaxy", BookPrice: 699})
	books = append(books, Books{BookID: "2", Book: "kool Book", BookPrice: 499})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/books", getAllbooks).Methods("GET")
	r.HandleFunc("/book/{id}", getOnebook).Methods("GET")
	r.HandleFunc("/book", createOnebook).Methods("POST")
	r.HandleFunc("/book/{id}", updateOnebook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteOnebook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllbooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(books)
}

func getOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one book")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for _, book := range books {
		if book.BookID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("No book found with given id")
	return
}

func createOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one book")
	w.Header().Set("Content-Type", "applicatioan/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)

	rand.Seed(time.Now().UnixNano())
	book.BookID = strconv.Itoa(rand.Intn(100))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return

}

func updateOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one book")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for index, book := range books {
		if book.BookID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Books
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.BookID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

func deleteOnebook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one book")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for index, book := range books {
		if book.BookID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}
