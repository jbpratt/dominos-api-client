package main

import (
	"bytes"
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

func getStoreCoupon(couponID, storeID string) string {
	URL := apiURL + "/store/" + storeID + "/coupon/" + couponID + "?lang=en"
	return request(URL)
}

func validateOrder(order []byte) string {
	URL := apiURL + "validate-order/"
	return postRequest(URL, order)
}

func postRequest(url string, data []byte) string {
	dominosClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp, err := dominosClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.Status
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
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return (string(body))
}

func main() {
	//	orderType := orderTypes["delivery"]
	//	cityRegionOrPostalCode := "Atlanta, GA, 30305"
	//	streetAddress := "3800 Northside Dr NW"
	//	res := getStoreNearAddress(orderType, cityRegionOrPostalCode, streetAddress)
	//	fmt.Println(res)
	//	fmt.Println("next response ==============================")
	//	storeID := "6342"
	//	res = getStoreInfo(storeID)
	//	fmt.Println(res)
	//	fmt.Println("next response ==============================")
	//res := getStoreMenu(storeID)
	//fmt.Println(res)
	//	res := getStoreCoupon("9193", storeID)
	res := validateOrder([]byte(`{"Order":{"Address":{"PostalCode":"90210"},"Coupons":[{"Code":"9193","Qty":1,"ID":3}],"CustomerID":"","Email":"","Extension":"","FirstName":"","LastName":"","LanguageCode":"en","OrderChannel":"OLO","OrderID":"FBgxwgEGQXG-g29uQEgC","OrderMethod":"Web","OrderTaker":null,"Payments":[],"Phone":"","PhonePrefix":"","Products":[],"ServiceMethod":"Carryout","SourceOrganizationURI":"order.dominos.com","StoreID":"7804","Tags":{},"Version":"1.0","NoCombine":true,"Partners":{"DOMINOS":{"Tags":{}}},"OrderInfoCollection":[]}}`))
	fmt.Println(res)
}

// curl for carryout
// curl 'https://order.dominos.com/power/store-locator?type=Carryout&c=35404&s=' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true' -H 'DNT: 1'

// curl for delivery
// curl 'https://order.dominos.com/power/store-locator?type=Delivery&c=ATLANTA%2C%20GA%2030305-1033&s=3800%20NORTHSIDE%20DR%20NW' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 30cfe3d96472371dd857c49b8e7b9e64212b8bd95ef6c630e41efd6ddc8b6ef5' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#b1b8d69f22774309a0ad651ed69c607d#1550286006' -H 'DNT: 1'

// curl 'https://order.dominos.com/power/store/5840/menu?lang=en&structured=true' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 3c6bced531cfe3d9647313f6db57caa07e11ebfa7eaa72f4f8142cf128679a31' -H 'Pragma: no-cache' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#8d00bb0ecb2544d9b64be0d26db1a06e#1550329931' -H 'DNT: 1'

// curl 'https://order.dominos.com/power/validate-order' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Content-Type: application/json; charset=utf-8' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: 06ac29669b31dc8b6ff4137983f90490ea0d69c78b673706e23e48a3528954a3' -H 'DNT: 1' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#62968c2114cd4a9fa214f6c7f031aba9#1550366132' --data '{"Order":{"Address":{"PostalCode":"90210"},"Coupons":[{"Code":"9193","Qty":1,"ID":3}],"CustomerID":"","Email":"","Extension":"","FirstName":"","LastName":"","LanguageCode":"en","OrderChannel":"OLO","OrderID":"FBgxwgEGQXG-g29uQEgC","OrderMethod":"Web","OrderTaker":null,"Payments":[],"Phone":"","PhonePrefix":"","Products":[],"ServiceMethod":"Carryout","SourceOrganizationURI":"order.dominos.com","StoreID":"7804","Tags":{},"Version":"1.0","NoCombine":true,"Partners":{"DOMINOS":{"Tags":{}}},"OrderInfoCollection":[]}}'

// curl 'https://order.dominos.com/power/validate-order' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Content-Type: application/json; charset=utf-8' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'X-DPZ-D: ae34112fc33844d2177de449aa016cda2f9a598e78bc547ef91e629b7a91bceb' -H 'DNT: 1' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#0bc8d264dcec43eb9aa63b20024c53fb#1550454022' --data '{"Order":{"Address":{"City":"ATLANTA","Region":"GA"},"Coupons":[{"Code":"9193","Qty":1,"ID":3}],"CustomerID":"","Email":"","Extension":"","FirstName":"","LastName":"","LanguageCode":"en","OrderChannel":"OLO","OrderID":"","OrderMethod":"Web","OrderTaker":null,"Payments":[],"Phone":"","PhonePrefix":"","Products":[],"ServiceMethod":"Carryout","SourceOrganizationURI":"order.dominos.com","StoreID":"5707","Tags":{},"Version":"1.0","NoCombine":true,"Partners":{"DOMINOS":{"Tags":{}}},"OrderInfoCollection":[]}}'

// curl 'https://order.dominos.com/power/price-order' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://order.dominos.com/en/assets/build/xdomain/proxy.html' -H 'Content-Type: application/json; charset=utf-8' -H 'Market: UNITED_STATES' -H 'DPZ-Language: en' -H 'DPZ-Market: UNITED_STATES' -H 'DPZ-Source: DSSPriceOrder' -H 'X-DPZ-D: ae34112fc33844d2177de449aa016cda2f9a598e78bc547ef91e629b7a91bceb' -H 'DNT: 1' -H 'Connection: keep-alive' -H 'Cookie: check=true; mbox=session#0bc8d264dcec43eb9aa63b20024c53fb#1550454022' --data '{"Order":{"Address":{"City":"ATLANTA","Region":"GA"},"Coupons":[{"Code":"9193","Qty":1,"ID":3}],"CustomerID":"","Email":"","Extension":"","FirstName":"","LastName":"","LanguageCode":"en","OrderChannel":"OLO","OrderID":"DiuE2ePMfJVke7_QSZjE","OrderMethod":"Web","OrderTaker":null,"Payments":[],"Phone":"","PhonePrefix":"","Products":[{"Code":"12SCREEN","Qty":1,"ID":4,"isNew":false,"Options":{"X":{"1/1":"1"},"C":{"1/1":"1"}}},{"Code":"CKRGCBT","Qty":1,"ID":5,"isNew":false,"Options":{"K":{"1/1":"1"},"Td":{"1/1":"1"}}},{"ID":0,"Code":"RANCH","Qty":1,"CategoryCode":"Sides","IsNew":true,"NeedsCustomization":false,"AutoRemove":false,"Fulfilled":false,"Tags":{}}],"ServiceMethod":"Carryout","SourceOrganizationURI":"order.dominos.com","StoreID":"5707","Tags":{},"Version":"1.0","NoCombine":true,"Partners":{"DOMINOS":{"Tags":{}}},"OrderInfoCollection":[]}}'
