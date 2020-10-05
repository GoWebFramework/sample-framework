package sample

import "net/http"

type handler func(w http.ResponseWriter, r *http.Request)

type Instance struct {
	Host   string
	Routes map[string]handler
}

// register handler
func (i *Instance) GET(path string, handler handler) {
	if _, exists := i.Routes[path]; exists {
		panic("duplicate path !, path already exist")
	}
	http.HandleFunc(path, handler)
}

func (i *Instance) Run() {
	http.ListenAndServe(i.Host, nil)
}
