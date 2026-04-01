package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")     // 印在 terminal
	fmt.Fprintln(w, "Hello, World!") // 回傳給 client (curl 或瀏覽器)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
