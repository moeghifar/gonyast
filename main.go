package main

import (
	"flag"
	"log"
	"strconv"
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
	var flagJeda = flag.Int("jeda", 60, "masukan detik jeda looping")
	var flagSaldo = flag.Float64("saldo", 0, "masukan saldo cryptocurrency saat ini")
	var flagLastBuy = flag.Float64("lastbuy", 0, "masukan harga beli terakhir anda")
	var flagCurrency = flag.String("currency", "", "masukan jenis currency (BTC / BCH / XZC)")
	var flagProfitThreshold = flag.Int64("profit", 0, "masukan threshold laba untuk notifikasi telegram")
	flag.Parse()
	// Write hello signature
	signature()
	if *flagSaldo == 0 || *flagLastBuy == 0 {
		log.Fatalln("Please insert the value to `-saldo`,`-lastbuy`,`-currency`,`-provit` flag!")
	}
	log.Println("jeda :", *flagJeda, " detik | saldo:", *flagSaldo, " | lastbuy :", strconv.FormatFloat(*flagLastBuy, 'f', 0, 64))
	log.Println("profit threshold :", *flagProfitThreshold, " | currency:", *flagCurrency)
	dataSet := btc.MyCash{
		Saldo:   *flagSaldo,
		LastBuy: *flagLastBuy,
	}
	dataConf := btc.BotConfig{
		Currency:        *flagCurrency,
		Sleep:           *flagJeda,
		ProfitThreshold: *flagProfitThreshold,
	}
	btc.Listen(dataSet, dataConf)
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
