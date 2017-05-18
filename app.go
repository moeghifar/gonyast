package main

import (
	"log"

	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/moeghifar/golang-workshop/src/util"
)

var (
	user  = "Ghiyast"
	pass  = "P4sswd"
	sPort = 1102
)

func init() {
	util.NewRedis("localhost:6379")
}

func main() {
	portString := fmt.Sprintf(":%d", sPort)
	// Write hello signature
	signature("1.0", portString)

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/user/", BasicAuth(User, user, pass))

	// Serving with http.ListenAndServe function which return fatal if error occured
	log.Fatal(http.ListenAndServe(portString, router))
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

// Hello ...
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	fmt.Fprintf(w, "Hella, Your name is %s", name)
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hella, welcome to NyastMach")
}

// User ...
func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "You're logged in to NyastMach")
}

func signature(version string, sPort string) {
	log.Println("/////////////")
	log.Println("\\ NyastMach")
	log.Println("\\ version", version)
	log.Println("\\ port", sPort)
	log.Println("/////////////")
}

func reverse(w string) (r string) {
	for _, v := range w {
		r = string(v) + r
	}
	return
}
