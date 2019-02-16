package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	orderType := "Delivery"
	cityStateZip := "Atlanta, GA, 30305"
	streetAddress := "3800 Northside Dr NW"

	//	URL := "https://order.dominos.com/power/store-locator?type=" + orderType + "&c=" + cityStateZip + "&s=" + streetAddress

	URL := "https://order.dominos.com/power/store-locator?"
	p := url.Values{"type": {orderType}, "c": {cityStateZip}, "s": {streetAddress}}
	URL = URL + p.Encode()
	dominosClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q-0.5")
	req.Header.Set("Market", "UNITED_STATES")
	req.Header.Set("DPZ-Language", "en")
	req.Header.Set("DPZ-Market", "UNITED_STATES")
	req.Header.Set("Connection", "keep-alive")

	res, getErr := dominosClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}

// curl for carryout
// curl 'https://order.dominos.com/power/store-locator?type=Carryout&c=35404&s=' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true' -H 'DNT: 1'

// curl for delivery
// curl 'https://order.dominos.com/power/store-locator?type=Delivery&c=ATLANTA%2C%20GA%2030305-1033&s=3800%20NORTHSIDE%20DR%20NW' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#b1b8d69f22774309a0ad651ed69c607d#1550286006' -H 'DNT: 1'
