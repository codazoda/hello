package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	port := getServerPort()
	http.HandleFunc("/", helloHandler)
	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Print(err.Error())
	}
}
