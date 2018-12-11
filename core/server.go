package core

// upload done
// convertion DONE
// SendEmail  DONE
// Remove Old file DONE
// Remove file after 5min TODO
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
	"strings"
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
	switch {
	case r.URL.Path == "/assets/sunlight.png":
		m.mutex.RLock()
		defer m.mutex.RUnlock()
		path := r.URL.Path[1:]
		data, _ := ioutil.ReadFile(string(path))
		w.Write(data)
	case r.URL.Path == "/":
		m.mutex.RLock()
		defer m.mutex.RUnlock()
		log.Dbg(m.conf.DBG, "Requesting ['%s'] with: ['%s']", r.URL.Path, r.Method)
		m.Index(w, r)
	case r.URL.Path == "/upload":
		if r.Method == "POST" {
			m.mutex.RLock()
			defer m.mutex.RUnlock()
			log.Dbg(m.conf.DBG, "Requesting ['%s'] with: ['%s']", r.URL.Path, r.Method)
			m.Upload(w, r)
		} else if r.Method == "GET" {
			io.WriteString(w, "GET is Unsuported method! in /upload use POST instead.")
		}
	case strings.Contains(r.URL.Path, "/files/"):
		log.Dbg(m.conf.DBG, "Requesting ['%s'] with ['%s']", r.URL.Path, r.Method)
		f := strings.Replace(r.URL.Path, "/files/", "", 1)
		if _, err := os.Stat("./tmp/" + f); os.IsNotExist(err) {
			io.WriteString(w, "<h1>No file with that name!</h1>")
		}
		m.HandleFileDownload(w, r, "./tmp/"+f)
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
	if err := os.Remove(path); err != nil {
		log.Err("Error while deleting the file %s %v", path, err)
	}
	m.Index(w, r)
}

func (m *Mux) HandleFileDownload(w http.ResponseWriter, r *http.Request, file string) {
	data, _ := ioutil.ReadFile(string(file))
	switch {
	case strings.HasSuffix(file, "jpg") || strings.HasSuffix(file, "jpeg") || strings.HasSuffix(file, "png"):
		r.Header.Add("Content-type", "image/*")
	case strings.HasSuffix(file, "pdf"):
		r.Header.Add("Content-type", "application/pdf")
	}
	w.Write(data)
}
