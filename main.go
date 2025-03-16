package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! the time is: %s", time.Now())
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 9090...")
	http.ListenAndServe(":9090", nil)
}
