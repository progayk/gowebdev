package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// check if any cookie exists in the request, if not create a new one
	c := getCookie(w, r)

	// process form submission
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			log.Println("form couldn't processed", err)
		}
		defer mf.Close()

		// file extension
		ext := strings.Split(fh.Filename, ".")[1]
		// gen new SHA1
		h := sha1.New()
		io.Copy(h, mf)
		// get hex value of SHA1(fh.filename)
		// concatenate file name and extension
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// get working dir
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		// join path
		path := filepath.Join(wd, "public", "pics", fname)
		// create a new file with concatenated name
		nf, err := os.Create(path)
		if err != nil {
			log.Fatalln("couldnt create the file", err)
		}
		defer nf.Close()
		// drag the cursor on mf to file start by using seek()
		mf.Seek(io.SeekStart, io.SeekStart)
		// copy the content into nf
		_, err = io.Copy(nf, mf)
		if err != nil {
			log.Println("coulndt write to the file", err)
		}
		// add filename to cookie values
		v := c.Value + "|" + fname
		c.Value = v
		http.SetCookie(w, c)
	}

	vx := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", vx)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("sID")
	if err != nil { // not nil if cookie doesn't exist
		// gen new ID
		uID, err := uuid.NewV4()
		if err != nil {
			log.Println("couldnt gen a an ID", err)
		}

		// create a cookie
		c = &http.Cookie{
			Name:  "sID",
			Value: uID.String(),
		}
		// set max age to 30 sec
		c.MaxAge = 30

		// write cookie to response
		http.SetCookie(w, c)
	}
	return c
}


