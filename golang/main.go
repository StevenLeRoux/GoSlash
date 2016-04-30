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
	"strconv"
	"strings"
)

// Config struct aims to offer a structure for loading configuration
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

type goslash struct {
	config Config
	store  store.Store
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: "+strconv.Itoa(status))
}

func (g goslash) GoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raw := params["alias"]
	args := strings.Split(raw, "/")
	fmt.Println("Args: ")
	fmt.Println(args)
	alias := args[0]
	values, ok := g.store.Get(alias)

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

func (g goslash) PutHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raw := params["alias"]
	args := strings.Split(raw, "/")
	fmt.Println("Args: ", args)
	alias := args[0]
	v := common.Values{
		Alias:       alias,
		Target:      "Target",
		User:        "User",
		Created:     "Created",
		Modified:    "Modified",
		Description: "Description",
	}
	err := g.store.Put(v)
	if err != nil {
		log.Println("PutHandler() - s.Put() => err")
		errorHandler(w, r, http.StatusNotModified)
		return
	}
	log.Println("no error from s.Put() in PutHandler")

	return
}

func Get200(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "200 OK")
}

func (g goslash) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	g.store.Update()
	w.WriteHeader(http.StatusAccepted)
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

	g := &goslash{
		config: cfg,
		store:  *s,
	}

	r := mux.NewRouter()
	r.HandleFunc("/200", Get200).Methods("GET")
	r.HandleFunc("/update", g.UpdateHandler).Methods("GET")
	r.HandleFunc("/set/{alias:.+}", g.PutHandler).Methods("PUT")
	r.HandleFunc("/{alias:.+}", g.GoHandler).Methods("GET")
	http.Handle("/", r)

	// HTTP - port 80
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}

}
