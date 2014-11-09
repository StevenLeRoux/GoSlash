package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/gorilla/mux"
	"le-roux.info/goslash/redirect"
	"le-roux.info/goslash/store"
	"log"
	"net/http"
	"strings"
)

type Config struct {
	Auth struct {
		Enabled string
	}
	Store struct {
		Engine   string
		Location string
	}
	Metrics struct {
		Backend string
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}

func GoHandler(w http.ResponseWriter, r *http.Request, s store.Store) {
	params := mux.Vars(r)
	raw := params["alias"]
	args := strings.Split(raw, "/")
	fmt.Println("Args: ")
	fmt.Println(args)
	alias := args[0]
	target := s.Get(alias).Target
	fmt.Println("target : " + target)
	if target == "" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	//http.Redirect(w, r, target, http.StatusFound)
	p := redirect.Apply(target, args)
	fmt.Println("RedirectTo : " + p)
	//if err != nil {
	http.Redirect(w, r, p, http.StatusFound)
	//} else {
	//	fmt.Println("error: " + err.Error())
	//	errorHandler(w, r, http.StatusNotFound)
	//}

	return
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "200 OK")
}

func main() {
	// flag

	var cfg Config
	var err error
	var s *store.Store
	err = gcfg.ReadFileInto(&cfg, "goslash.cfg")

	s, err = store.New(cfg.Store.Engine, cfg.Store.Location)

	r := mux.NewRouter()
	r.HandleFunc("/200", TestHandler).Methods("GET")
	r.HandleFunc("/{alias:.+}", func(w http.ResponseWriter, r *http.Request) {
		GoHandler(w, r, *s)
	}).Methods("GET")
	http.Handle("/", r)

	// HTTP - port 80
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}

}
