package main

import (
	"fmt"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request)  {
	//io.WriteString(w, "This is a test")
	fmt.Fprintf(w, "HelloWord")
}

func main()  {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8080", nil)
}
