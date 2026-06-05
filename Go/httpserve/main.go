package main

import (
	"fmt";
	"net/http"
)

func index (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello Web")
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/",index)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
