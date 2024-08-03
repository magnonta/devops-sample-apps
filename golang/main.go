package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
    pwd, _ := os.Getwd()
    fmt.Printf("Current working directory: %s\n", pwd)

    p12exists := Exists(pwd + "/file.p12")
    if !p12exists {
        fmt.Println(pwd + "/file.p12 not found");
        return;
    }

    fmt.Println("file.p12 found, starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}