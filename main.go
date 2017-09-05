package main

import (
	"log"
	"time"

	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/moeghifar/gonyast/src/sekolah"
	"github.com/moeghifar/gonyast/src/util"
)

var (
	appname = "Gonyast"
	user    = "Ghiyast"
	pass    = "P4sswd"
	sPort   = util.Config.Port
	version = "1.0"
)

func init() {
	startTime := time.Now()
	if err := util.Init(); err != nil {
		log.Fatal("[ERROR]", err)
	}
	startTime = time.Now()
	log.Println("[LOG] initiate util package done in", fmt.Sprintf("%f", time.Since(startTime).Seconds()*1000), "ms")
	if err := sekolah.Init(); err != nil {
		log.Fatal("[ERROR]", err)
	}
	log.Println("[LOG] initiate sekolah package done in", fmt.Sprintf("%f", time.Since(startTime).Seconds()*1000), "ms")
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
	// API sekolah
	router.GET("/api/get/sekolah/v1", sekolah.GetSekolah)
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
	log.Println("//////////////////")
	log.Println("\\", appname, " v", version)
	log.Println("\\ port", sPort)
	log.Println("//////////////////")
}

func reverse(w string) (r string) {
	for _, v := range w {
		r = string(v) + r
	}
	return
}
