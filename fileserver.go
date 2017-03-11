package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// A static web server serving a current path
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [path]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Using current dir as path\n")
		fmt.Println("Server starting on Port 8080")
		log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	}
	fmt.Println("Server starting on Port 8080")
	fmt.Println("Server Hosting: ", os.Args[1])
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Args[1]))))
}
