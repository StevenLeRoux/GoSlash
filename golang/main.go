package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/gorilla/mux"
	"le-roux.info/goslash/golang/redirect"
	"le-roux.info/goslash/golang/store"

	"le-roux.info/goslash/golang/store/common"
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
	values, ok := s.Get(alias)

	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	target := values.Target

	fmt.Println("target : " + target)
	//http.Redirect(w, r, target, http.StatusFound)
	p := redirect.Apply(target, args)
	log.Println("RedirectTo : ", p)
	//if err != nil  {
	http.Redirect(w, r, p, http.StatusFound)
	//} else {
	//	fmt.Println("error: " + err.Error())
	//	errorHandler(w, r, http.StatusNotFound)
	//}

	return
}

func PutHandler(w http.ResponseWriter, r *http.Request, s store.Store) {
	params := mux.Vars(r)
	raw := params["alias"]
	args := strings.Split(raw, "/")
	fmt.Println("Args: ", args)
	alias := args[0]
	v := common.Values{Alias: alias, Target: "Target", User: "User", Created: "Created", Modified: "Modified", Description: "Description"}
	err := s.Put(v)
	if err != nil {
		log.Println("PutHandler() - s.Put() => err")
		errorHandler(w, r, http.StatusNotModified)
		return
	}
	log.Println("no error from s.Put() in PutHandler")

	return
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "200 OK")
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, s store.Store) {
	s.Update()
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "201 Accepted")
}

func main() {
	// flag

	var cfg Config
	var s *store.Store
	err := gcfg.ReadFileInto(&cfg, "goslash.cfg")

	s, err = store.New(cfg.Store.Engine, cfg.Store.Location)
	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/200", TestHandler).Methods("GET")
	r.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		UpdateHandler(w, r, *s)
	}).Methods("GET")
	r.HandleFunc("/set/{alias:.+}", func(w http.ResponseWriter, r *http.Request) {
		PutHandler(w, r, *s)
	}).Methods("PUT")
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
