package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Category string  `json:"category"`
}

var list []Product = []Product{
	{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"},
	{Id: 101, Name: "Pencil", Cost: 5, Category: "Stationary"},
	{Id: 102, Name: "Marker", Cost: 50, Category: "Stationary"},
	{Id: 103, Name: "Mouse", Cost: 250, Category: "Electronics"},
}

func getProductById(id int) *Product {
	for _, p := range list {
		if id == p.Id {
			return &p
		}
	}
	return nil
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "All the products will be served")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	if id, err := strconv.Atoi(idString); err == nil {
		product := getProductById(id)
		if product != nil {
			if err := json.NewEncoder(w).Encode(product); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "bad request error", http.StatusBadRequest)
	}
}

func addNewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "bad request error", http.StatusBadRequest)
		return
	}
	list = append(list, newProduct)
	w.WriteHeader(http.StatusCreated)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", getAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", getOneProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", addNewProduct).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}
