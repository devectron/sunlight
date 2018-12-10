package core

import (
	"crypto/md5"
	//"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/devectron/sunlight/log"
)

// Server interface.
type Server interface {
	Index(w http.ResponseWriter, r *http.Request)
	Convertor(w http.ResponseWriter, r *http.Request)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// Mux mutex.
type Mux struct {
	Server
	mutex sync.RWMutex
	conf  Config
	data  SiteData
}

// SiteData data of the site.
type SiteData struct {
	Title     string
	ErrorBool bool
	Error     string
	NbrConv   int
	Users     string
	Token     string
}

// StartListening listen to a given port.
// upload done
// convertion TODO
// SendEmail  TODO
// Remove Old file TODO
// Remove Email after 5min TODO
func StartListening(c Config) {
	s := SiteData{
		Title:   "Sunlight | Documents Convertor",
		NbrConv: 1002,
		Users:   "900 Users",
	}
	m := &Mux{
		conf: c,
		data: s,
	}
	log.Inf("Listening on: localhost:%s", c.ServerPort)
	err := http.ListenAndServe(":"+c.ServerPort, m)
	if err != nil {
		log.Err("%v", err)
	}
}

// ServeHTTP http handler.
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/assets/sunlight.png":
		m.mutex.RLock()
		defer m.mutex.RUnlock()
		path := r.URL.Path[1:]
		data, _ := ioutil.ReadFile(string(path))
		w.Write(data)
	case "/":
		m.mutex.RLock()
		defer m.mutex.RUnlock()
		log.Dbg(m.conf.DBG, "Requesting ['%s'] with: ['%s']", r.URL.Path, r.Method)
		m.Index(w, r)
	case "/upload":
		if r.Method == "POST" {
			m.mutex.RLock()
			defer m.mutex.RUnlock()
			log.Dbg(m.conf.DBG, "Requesting ['%s'] with: ['%s']", r.URL.Path, r.Method)
			m.Upload(w, r)
		} else if r.Method == "GET" {
			io.WriteString(w, "GET is Unsuported method! in /upload")
		}
	default:
		m.mutex.RLock()
		defer m.mutex.RUnlock()
		log.Dbg(m.conf.DBG, "Getting unsuported path [ '%s' ] [ '%s' ]", r.URL.Path, r.Method)
		log.War("Redirecting to [ '/' ].")
		http.Redirect(w, r, "/", 302)
	}
}

// Index return index page.
func (m *Mux) Index(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.New("index.html").Parse(INDEX)
	if err != nil {
		log.Err("Error html parser %v", err)
	}
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	m.data.Token = token
	m.data.Error = "Test"
	m.data.ErrorBool = false
	htmlTemplate.Execute(w, m.data)
}

// Upload upload file.
func (m *Mux) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) //memory storage
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Err("Error While uploading file %v", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Err("Error while reading data %v", err)
	}
	path := "./tmp/" + handler.Filename
	err = ioutil.WriteFile(path, data, 0666)
	log.Inf("Uploading file %s lenght:%d", handler.Filename, handler.Size)
	if err != nil {
		log.Err("Error while writing to the file %v", err)
	}
	format := r.Form["type"]
	log.War("Converting File ...")
	dstfile, err := Convertor(path, m.conf.ConvertApi, format[0])
	if err != nil {
		log.Err("Error while converting file %v", err)
	}
	log.Inf("Sending email ...")
	email := r.PostFormValue("email")
	SendMail(email, dstfile, m.conf.MailApiPublic, m.conf.MailApiPrivate)
	log.War("Removing file ...")
	//	err := os.Remove(path)
	if err := os.Remove(path); err != nil {
		log.Err("Error while deleting the file %s %v", path, err)
	}
	m.Index(w, r)
	//http.Redirect(w, r, "/", 200)
}
