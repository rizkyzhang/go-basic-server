package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server starting at port 8080\n")
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api", rootHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/api")	{
		fmt.Fprintln(w, "404 not found", http.StatusNotFound)
		return
	}

	if (r.Method != "GET") {
		fmt.Fprintln(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	
	fmt.Fprintf(w, "Welcome to basic go server!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "age: %s\n", age)
	fmt.Fprintf(w, "address: %s", address)
}