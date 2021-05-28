package Service

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func telegram() {
	queryString := ""
	sendTelegram("val", queryString)
}

func sendTelegram(channelID, queryString string) {
	telegram := url.QueryEscape(queryString)
	_ = time.Now()
	resp, err := http.Get("https://api.telegram.org/bot1870108950:AAFJ9uBcBwCMoaunX8XtoAXFuAtynxEbXHw/sendMessage?chat_id=" + channelID + "&text=" + telegram)
	_ = time.Now()
	if err != nil || resp.StatusCode != 200 {
		if resp != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("No idea failed", err)
			} else {
				log.Println("telegram failed", body)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println("successfully sent to telegram channel", telegram)
	}
}
