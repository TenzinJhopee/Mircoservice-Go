package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	fmt.Println("hiii")

	http.HandleFunc()

	http.ListenAndServe(":9090", nil)
}
