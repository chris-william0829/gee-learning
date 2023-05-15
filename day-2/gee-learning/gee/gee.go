package gee

import(
	"net/http"
)

//request handler
type HandleFunc func(*Context)

//Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

//New is the constructor of gee.Engine
func New() *Engine{
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc){
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, req)
	engine.router.handle(c)
}