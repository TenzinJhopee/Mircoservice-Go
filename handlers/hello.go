package handlers


import {
	"net/http"
    "log"
	"io/ioutil"
	"fmt"
}

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{}
}

func (h*Hello) ServeHttp(rw http.ResponseWriter, r *http.Request0) {
		h.l.Println("Hello World")

		log.Println("Hello WOrld")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return 
		}
		fmt.Fprintf(rw, "Hello %s", d)
}