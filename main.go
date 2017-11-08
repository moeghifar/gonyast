package main

import (
	"log"
	"time"

	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/moeghifar/gonyast/src/btc"
	"github.com/moeghifar/gonyast/src/sekolah"
	"github.com/moeghifar/gonyast/src/util"
)

var (
	appname = "Gonyast"
	sPort   = util.Config.Port
	version = "1.0"
)

func init() {
	startTime := time.Now()
	if err := util.Init(); err != nil {
		log.Fatal("[ERROR]", err)
	}
	log.Println("[LOG] initiate Util package done in", fmt.Sprintf("%f", time.Since(startTime).Seconds()*1000), "ms")
}

func main() {
	// Write hello signature
	signature()
	btc.TelegramListener()
}

func router(router *httprouter.Router) {
	portString := fmt.Sprintf(":%s", sPort)

	// List of router
	router.GET("/", Index)
	// API sekolah
	router.GET("/api/get/sekolah/v1", sekolah.GetSekolah)

	// Serving with http.ListenAndServe function which return fatal if error occured
	log.Fatal(http.ListenAndServe(portString, router))
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hella, welcome to %s", appname)
}

func signature() {
	log.Println("//////////////////")
	log.Println("\\", appname, " v", version)
	log.Println("\\ port", sPort)
	log.Println("//////////////////")
}
