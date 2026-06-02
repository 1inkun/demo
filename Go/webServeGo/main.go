package main

import (
	"fmt"
	"log"
	"net/http"
)

func formmethod(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() has error: %s \n",err)
		return
	}
	fmt.Fprintln(w,"POST Request success")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"name is %s\n",name)
	fmt.Fprintf(w,"address is %s\n",address)
}

func hellomethod(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Not support method",http.StatusBadRequest)
		return
	}
	w.Write([]byte("HelloWeb"))
}


func main(){
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",hellomethod)
	http.HandleFunc("/form",formmethod)



	log.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
