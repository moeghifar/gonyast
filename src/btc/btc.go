package btc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
)

// GetBtc to get btc info every 30s
func GetBtc() {
	apiURL := "https://vip.bitcoin.co.id/api/btc_idr/webdata"
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
	log.Println(fmt.Sprintf("bitcoin last price at %s (bitcoin.co.id) : %s", time.Now().Format(time.RFC822), resultBody.LP))
}
