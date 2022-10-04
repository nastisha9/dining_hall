package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var orders []Order

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func postOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(&order)

	time.Sleep(time.Second * 3)

	ord, err := PrettyStruct(order)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nDining-hall recieved the order:\n", ord)
}

func main() {
	router := mux.NewRouter()

	//URL path and the function to handle
	router.HandleFunc("/distribution", postOrder).Methods("POST")

	initTables()
	waitToOrder()

	http.ListenAndServe(":3030", router)
}
