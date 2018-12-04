package core

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/devectron/sunlight/log"
)

type Mux struct {
	server Server
	mutex  sync.RWMutex
}

type config struct {
	title string
}

type Server interface {
	index(w http.ResponseWriter, r *http.Request)
	convertor(w http.ResponseWriter, r *http.Request)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func StartListening() {
	m := &Mux{}
	http.HandleFunc("/", m.index)
	http.HandleFunc("/doconv", m.convertor)
	log.Inf("Listening on: localhost:7375")
	err := http.ListenAndServe(":7375", m)
	if err != nil {
		log.Err("%v", err)
	}
}
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		m.index(w, r)
	} else if r.URL.Path == "/doconv" {
		m.convertor(w, r)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func (m *Mux) index(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.New("index.html").Parse(INDEX)
	if err != nil {
		log.Err("Error html parser %v", err)
	}
	htmlTemplate.Execute(w, nil)
}
func (m *Mux) convertor(w http.ResponseWriter, r *http.Request) {

}
