package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var apiURL = "https://order.dominos.com/power"
var orderTypes = map[string]string{"delivery": "delivery", "carryout": "carryout"}

func getStoreNearAddress(orderType, cityRegionOrPostalCode, streetAddress string) string {
	p := url.Values{"type": {orderType}, "c": {cityRegionOrPostalCode}, "s": {streetAddress}}
	URL := apiURL + "/store-locator?" + p.Encode()
	return request(URL)
}

func getStoreInfo(storeID string) string {
	URL := apiURL + "/store/" + storeID + "/profile"
	return request(URL)
}

func getStoreMenu(storeID string) string {
	URL := apiURL + "/store/" + storeID + "/menu?lang=en&structured=true"
	return request(URL)
}

func request(url string) string {
	dominosClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")

	res, getErr := dominosClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return (string(body))
}

func main() {
	orderType := orderTypes["delivery"]
	cityRegionOrPostalCode := "Atlanta, GA, 30305"
	streetAddress := "3800 Northside Dr NW"
	res := getStoreNearAddress(orderType, cityRegionOrPostalCode, streetAddress)
	fmt.Println(res)
	fmt.Println("next response ==============================")
	storeID := "6342"
	res = getStoreInfo(storeID)
	fmt.Println(res)
	fmt.Println("next response ==============================")
	res = getStoreMenu(storeID)
	fmt.Println(res)

}

// curl for carryout
// curl 'https://order.dominos.com/power/store-locator?type=Carryout&c=35404&s=' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true' -H 'DNT: 1'

// curl for delivery
// curl 'https://order.dominos.com/power/store-locator?type=Delivery&c=ATLANTA%2C%20GA%2030305-1033&s=3800%20NORTHSIDE%20DR%20NW' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#b1b8d69f22774309a0ad651ed69c607d#1550286006' -H 'DNT: 1'
