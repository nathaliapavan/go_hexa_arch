package server

import (
	"github.com/nathaliapavan/go_hexa_arch/application"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"os"
	"time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)
	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}