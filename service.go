package main

//go:generate go run scripts/ui.go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/StevenLeRoux/goslash/golang/redirect"
	"github.com/StevenLeRoux/goslash/golang/store"
	"github.com/StevenLeRoux/goslash/golang/store/model"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	log "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

type goslash struct {
	config *viper.Viper
	store  store.Store
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(status)
	}
	fmt.Fprintf(w, "Error: "+strconv.Itoa(status))
}

func (g goslash) GoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raw := params["alias"]
	args := strings.Split(raw, "/")
	fmt.Println("Args: ")
	fmt.Println(args)
	alias := args[0]
	value, ok := g.store.Get(alias)
	fmt.Println("GoHandler status : ", ok)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	target := value.Target

	fmt.Println("target : " + target)
	//http.Redirect(w, r, target, http.StatusFound)
	p := redirect.Apply(target, args)
	log.INFO.Println("RedirectTo : ", p)
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
	v := model.Values{
		Alias:       alias,
		Target:      "Target",
		User:        "User",
		Created:     "Created",
		Modified:    "Modified",
		Description: "Description",
	}
	err := g.store.Put(v)
	if err != nil {
		log.ERROR.Println("Error while trying to Put() on store =>", err)
		errorHandler(w, r, http.StatusNotModified)
		return
	}

	return
}

func Get200(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "200 OK")
}

func GetAlias(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, alias)
}

func (g goslash) ReloadHandler(w http.ResponseWriter, r *http.Request) {
	g.store.Reload()
	w.WriteHeader(http.StatusAccepted)
}

func (g goslash) DumpHandler(w http.ResponseWriter, r *http.Request) {
	values, ok := g.store.Dump()
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data, err := json.Marshal(values)
	if err != nil {

		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func runService(cmd *cobra.Command) {

	config, err := InitializeConfig()
	if err != nil {
		log.FATAL.Fatalln(err)
	}

	var s *store.Store

	s, err = store.New(config)
	if err != nil {
		log.ERROR.Panic(err)
	}

	g := &goslash{
		config: config,
		store:  *s,
	}

	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/200", Get200).Methods("GET")
	r.HandleFunc("/reload", g.ReloadHandler).Methods("GET")
	r.HandleFunc("/dump", g.DumpHandler).Methods("GET")
	r.HandleFunc("/set/{alias:.+}", g.PutHandler).Methods("PUT")
	r.HandleFunc("/{alias:.+}", g.GoHandler).Methods("GET")
	r.HandleFunc("/", GetAlias).Methods("GET")
	http.Handle("/", r)
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.FATAL.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}

}
