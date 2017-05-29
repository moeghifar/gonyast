package main

import (
	"log"

	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/moeghifar/nyastmach/src/sklh"
	"github.com/moeghifar/nyastmach/src/util"
)

var (
	appname = "NyastMach"
	user    = "Ghiyast"
	pass    = "P4sswd"
	sPort   = 1102
	version = "1.0"
)

func init() {
	util.NewRedis("localhost:6379")
	util.InitDatabase()
	err := sklh.Init()
	if err != nil {
		log.Fatal("[FATAL] Failed initiating sklh package ->", err)
	}
}

func main() {
	// Write hello signature
	signature()
	// write router and port listening
	router(httprouter.New())
}

// BasicAuth ...
func BasicAuth(h httprouter.Handle, requiredUser, requiredPass string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		user, password, hasAuth := r.BasicAuth()
		if hasAuth && user == requiredUser && password == requiredPass {
			h(w, r, p)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func router(router *httprouter.Router) {
	portString := fmt.Sprintf(":%d", sPort)

	// List of router
	router.GET("/", Index)
	// API sklh
	router.GET("/api/v1/get_sklh", sklh.GetSklh)
	// test response with auth
	router.GET("/user/", BasicAuth(User, user, pass))

	// Serving with http.ListenAndServe function which return fatal if error occured
	log.Fatal(http.ListenAndServe(portString, router))
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hella, welcome to %s", appname)
}

// User ...
func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "You're logged in to %s", appname)
}

func signature() {
	log.Println("///////////////")
	log.Println("\\ ", appname, " \\")
	log.Println("\\ version", version, "\\")
	log.Println("\\  port", sPort, " \\")
	log.Println("///////////////")
}

func reverse(w string) (r string) {
	for _, v := range w {
		r = string(v) + r
	}
	return
}
