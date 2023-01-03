package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Set up a file server to serve static HTML files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the server
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
