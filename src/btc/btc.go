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
	// MyCash ...
	MyCash struct {
		LastMoney float64
		Saldo     float64
		LastBuy   float64
	}
	// BotConfig ...
	BotConfig struct {
		Currency        string
		ProfitThreshold int64
		Sleep           int
	}
	telegramBot struct {
		APIKey string
		APIURL string
	}
)

// Listen ...
func Listen(dataSet MyCash, dataConf BotConfig) {
	for {
		GetBtc(dataSet, dataConf)
		time.Sleep(time.Duration(dataConf.Sleep) * time.Second)
	}
}

// GetBtc to get btc info through http request
func GetBtc(dataSet MyCash, dataConf BotConfig) (info string) {
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
	log.Println("Response Code:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	var resultBody respBody
	json.Unmarshal(body, &resultBody)

	lastPriceInt, _ := strconv.ParseInt(resultBody.LP, 10, 64)

	// get redis
	prevPrice, err := util.GetRedis("last_price")
	prevSellProfit, err := util.GetRedis("sell_profit")
	var msg string
	// var changed bool
	if prevPrice != "" {
		prevPriceInt, _ := strconv.ParseInt(prevPrice, 10, 64)
		if prevPriceInt == lastPriceInt {
			msg = "Not Changing"
		} else if prevPriceInt > lastPriceInt {
			// changed = true
			msg = "Getting Down"
		} else {
			// changed = true
			msg = "Getting Up"
		}
	}
	// set redis
	util.SetRedis(resultBody.LP, "last_price")

	fl, _ := strconv.ParseFloat(strconv.FormatFloat(float64(lastPriceInt), 'f', 0, 64), 64)
	conversion := dataSet.Saldo * fl
	lastBuyFormated, _ := strconv.ParseFloat(strconv.FormatFloat(dataSet.LastBuy, 'f', 0, 64), 64)
	profit := conversion - (lastBuyFormated * dataSet.Saldo)
	lastMoney := strconv.FormatFloat((lastBuyFormated * dataSet.Saldo), 'f', 0, 64)
	timeAt := time.Now().Format(time.RFC822)

	log.Println(fmt.Sprintf("%s [%s from previous]", dataConf.Currency, msg))
	log.Println(fmt.Sprintf("Check at %s", timeAt))
	log.Println(fmt.Sprintf("[LAST PRICE] Rp %s", resultBody.LP))
	log.Println(fmt.Sprintf("[LAST MONEY] Rp %s", lastMoney))
	log.Println(fmt.Sprintf("[PROFITS???] Rp %s", strconv.FormatFloat(profit, 'f', 0, 64)))

	intPrevSellProfit, _ := strconv.ParseInt(prevSellProfit, 10, 64)
	if int64(profit) > dataConf.ProfitThreshold && int64(profit) != intPrevSellProfit {
		messageToTelegram := fmt.Sprintf("your provit Rp %s", strconv.FormatFloat(profit, 'f', 0, 64))
		// set redis
		util.SetRedis(resultBody.LP, "sell_profit")
		SendTelegram(messageToTelegram)
	}
	return
}

// SendTelegram ...
func SendTelegram(getInfo string) {
	botConfig := &telegramBot{
		APIKey: "492963574:AAFXOign3WASxhyy_3obG1yebSP3qqQJ--Y",
		APIURL: "https://api.telegram.org/bot",
	}
	req, _ := http.PostForm(botConfig.APIURL+botConfig.APIKey+"/sendMessage", url.Values{"chat_id": {"@moeghifar_notify"}, "text": {getInfo}})
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
