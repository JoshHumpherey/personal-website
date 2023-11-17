package main

import (
	"fmt"
	"net/http"
)

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", landingPageHandler)

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
