package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)


}

func helloHandler(w http.ResponseWriter, r *http.Request){
	
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "404 Not Implemented", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello World!")
}



func main() {
	fileServer := http.FileServer(http.Dir("./static")) // Telling go lang to check static folder and it automatically search for index.html by default
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8089")

	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}

}