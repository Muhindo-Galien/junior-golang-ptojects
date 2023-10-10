package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() %v",err)
	}
	fmt.Fprintf(w, "Form submitted successfully")
	name :=r.FormValue("name")
	address :=r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address = %s\n",address)
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"404 not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main(){
	fileServe := http.FileServer(http.Dir("./static"))
	// Handle routes
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting on port 8080")

	if err:= http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
	}
}