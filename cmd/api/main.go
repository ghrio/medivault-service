package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		fmt.Println(userAgent)
		w.Write([]byte("Hello, World!"))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Server is listening on Port 8080")
	server.ListenAndServe()

}
