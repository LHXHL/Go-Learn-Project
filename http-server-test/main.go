package main

import (
	"fmt"
	"log"
	"net/http"
)

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)

	}
}

func main() {
	n := 5
	fmt.Println(fibonacci(n))
	fileServer := http.FileServer(http.Dir("www"))
	http.Handle("/", fileServer)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Listening on port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Fatal(err)
	}

	username := request.FormValue("username")
	password := request.FormValue("password")

	_, _ = fmt.Fprintf(writer, "Hello %s!\n", username)
	_, _ = fmt.Fprintf(writer, "U pwd is %s", password)
}
