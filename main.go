package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "The name is %s\n", name)
	fmt.Fprintf(w, "The address is %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form.html", formHandler)

	fmt.Println("Starting the server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/* Adding this stupid comment to check if
waka time api key connected
successfully
*/
