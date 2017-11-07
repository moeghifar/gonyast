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
	// log.Println(fmt.Sprintf("bitcoin cash last price at %s (bitcoin.co.id) : %s", time.Now().Format(time.RFC822), resultBody.LP))
	info = fmt.Sprintf("bitcoin cash last price at %s (bitcoin.co.id) : %s", time.Now().Format(time.RFC822), resultBody.LP)
	if (lastPriceInt - dataSet.Buy) > sellMinimum {
		percentSell := (currentSellPrice * 100) / dataSet.Limit
		// log.Println(fmt.Sprintf("jual sekarang untung %d percent yakni : %d", percentSell, currentSellPrice))
		info += fmt.Sprintf("\n---\njual sekarang untung %d percent yakni : %d", percentSell, currentSellPrice)
	}
	return
}

// TelegramListener ...
func TelegramListener() {
	botConfig := &telegramBot{
		APIKey: "492963574:AAEhdUk4WnZaAmZdtA1m5LMcIreQR1Ee4Ag",
		APIURL: "https://api.telegram.org/bot",
	}

	for {
		getInfo := GetBtc()
		log.Println(getInfo)
		req, _ := http.PostForm(botConfig.APIURL+botConfig.APIKey+"/sendMessage", url.Values{"chat_id": {"@mgf_notifier"}, "text": {getInfo}})
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		time.Sleep(5 * time.Second)
	}

}
