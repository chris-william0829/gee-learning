package gee

import(
	"fmt"
	"net/http"
)

//request handler
type HandleFunc func(http.ResponseWriter, *http.Request)

//Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

//New is the constructor of gee.Engine
func New() *Engine{
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc){
	key := method + "-" + pattern
	engine.router[key] = handler
}

//
func (engine *Engine) GET(pattern string, handler HandleFunc){
	engine.addRoute("GET", pattern, handler)
}

//
func (engine *Engine) POST(pattern string, handler HandleFunc){
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string)(err error){
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok{
		handler(w, req)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}