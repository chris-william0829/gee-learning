package main

import(
	"fmt"
	"net/http"

	"gee"
)

func main(){
	server := gee.New()
	server.GET("/",indexHandler)
	server.GET("/hello", helloHandler)
	server.RUN(":8080")
}

func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request){
	for k, v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}