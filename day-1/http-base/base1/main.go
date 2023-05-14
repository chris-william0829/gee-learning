package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	//bind handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	
	//listen localhost:8080 and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request){
	for k, v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}