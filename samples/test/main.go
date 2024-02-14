package main

import (
        "fmt"
        "net/http"
)

func main() {
        fmt.Println("test is OK.")
	resp, err := http.Get("http://www.google.com")
        if err != nil {
                panic(err)
        }
        fmt.Println(resp.Status)
}
