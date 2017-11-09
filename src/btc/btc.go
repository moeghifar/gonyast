package btc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/moeghifar/gonyast/src/util"
)

type (
	reqBody struct {
		CSRFToken string `json:"csrf_token"`
		Lang      string `json:"lang"`
	}
	respBody struct {
		LP string `json:"last_price"`
	}
	myCash struct {
		Saldo float64
		Buy   int64
		Limit int64
	}
	telegramBot struct {
		APIKey string
		APIURL string
	}
)

// TelegramListener ...
func TelegramListener() {
	for {
		GetBtc()
		time.Sleep(60 * time.Second)
	}
}

// GetBtc to get btc info through http request
func GetBtc() (info string) {
	apiURL := "https://vip.bitcoin.co.id/api/webdata/BCHIDR"
	data := url.Values{}
	data.Set("csrf_token", "5fbdec721a818fc0ec9ea2de62d936b60f57882d2830c4dce5f2a22a164fcd5e")
	data.Add("lang", "indonesia")
	u, _ := url.ParseRequestURI(apiURL)
	urlStr := u.String()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("==========================================================================")
	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	var resultBody respBody
	json.Unmarshal(body, &resultBody)

	dataSet := &myCash{
		Saldo: 0.26962776,
		Buy:   8500000,
		Limit: 20,
	}
	lastPriceInt, _ := strconv.ParseInt(resultBody.LP, 10, 64)
	sellMinimum := ((dataSet.Buy * dataSet.Limit) / 100)
	currentSellPrice := dataSet.Limit * lastPriceInt
	// get redis
	prevPrice, err := util.GetRedis("bitcoin_cash_cache")
	var msg string
	var changed bool
	if prevPrice != "" {
		prevPriceInt, _ := strconv.ParseInt(prevPrice, 10, 64)
		if prevPriceInt == lastPriceInt {
			msg = "Not Changing"
		} else if prevPriceInt > lastPriceInt {
			changed = true
			msg = "Getting Down"
		} else {
			changed = true
			msg = "Getting Up"
		}
	}
	util.SetRedis(resultBody.LP, "bitcoin_cash_cache")
	info = fmt.Sprintf("Last Price Rp %s [%s from previous]", resultBody.LP, msg)
	if (lastPriceInt - dataSet.Buy) > sellMinimum {
		percentSell := (currentSellPrice * 100) / dataSet.Limit
		info += fmt.Sprintf("\n---\njual sekarang untung %d percent yakni : %d", percentSell, currentSellPrice)
	}
	if changed {
		sendToTelegram(info)
	}
	log.Println(info)
	return
}

func sendToTelegram(getInfo string) {
	botConfig := &telegramBot{
		APIKey: "492963574:AAEhdUk4WnZaAmZdtA1m5LMcIreQR1Ee4Ag",
		APIURL: "https://api.telegram.org/bot",
	}
	req, _ := http.PostForm(botConfig.APIURL+botConfig.APIKey+"/sendMessage", url.Values{"chat_id": {"@mgf_notifier"}, "text": {getInfo}})
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
