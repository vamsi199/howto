package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type custm struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Custms []custm

func main() {
	//add records to structs
	Custms = append(Custms, custm{1, "Bobby"})
	Custms = append(Custms, custm{2, "Robin"})

	router := mux.NewRouter()
	//customer api
	router.HandleFunc("/customers", getCustomers).Methods("GET")                //gets all customers
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")            //gets one customer
	router.HandleFunc("/customers/{id}/{name}", editCustomer).Methods("PUT")    //gets one customer
	router.HandleFunc("/customers/{id}/{name}", createCustomer).Methods("POST") //creates new customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")      //deletes a customer

	log.Fatal(http.ListenAndServe(":8080", router))
}
func jMarshal(w http.ResponseWriter, r *http.Request, res []custm) {
	jResult, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "marshal Error...:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jResult)
}
func getCustomers(w http.ResponseWriter, r *http.Request) {
	//fetch all customers
	jResult, err := json.Marshal(Custms)
	if err != nil {
		http.Error(w, "marhsal error", http.StatusInternalServerError)
	}
	w.Write(jResult)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jResult := []custm{} //need to empty the slice before appending query results
	for _, c := range Custms {
		if strconv.Itoa(c.Id) == vars["id"] {
			jResult = append(jResult, c)
		}
	}
	jMarshal(w, r, jResult) //gives key and values whereas fmt.Fprint gives only values
}
func editCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "type conversion error...")
		return
	}
	for i, j := range Custms {
		if j.Id == nId {
			Custms[i].Name = vars["name"]
		}
	}
	jMarshal(w, r, Custms)

}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "type conversion error", http.StatusInternalServerError)
	}

	for _, i := range Custms {
		if i.Id == nId {
			http.Error(w, "customer already exists", http.StatusConflict)
			return
		}
	}
	var cus custm
	cus.Id = nId
	cus.Name = vars["name"] //did this by split(/) of vars
	Custms = append(Custms, cus)
	jMarshal(w, r, Custms)
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "conversion error", http.StatusInternalServerError)
	}
	for i, j := range Custms { //need how this worked ?
		if j.Id == nId {
			Custms = append(Custms[:i], Custms[i+1:]...)
			//delete(Custms,nId)//helpful only if custms is a map of slice
			break
		}
	}
	jMarshal(w, r, Custms)
}