package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type invoice struct {
	InvoiceId string
	Name string
	StreetAndNumber string
	Zip string
	City string
	Country string
	Date string
	DueDate string
}

func checkErr(err error){
	if err != nil {panic(err)}
}

func main(){
	fmt.Print("PDF reporting service started.\n")
	http.HandleFunc("/", getInvoice)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func getInvoice(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		w.WriteHeader(400)
		w.Write([]byte("Wrong HTTP-method retard :)"))
		return
	}
	var inv invoice
	json.NewDecoder(r.Body).Decode(&inv)
	w.Header().Add("Content-Type", "application/octet-stream")
	createInvoice(inv, w)
}