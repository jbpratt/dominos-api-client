package dominos

// Address represents the address of a user
type Address struct {
	Street       string `json:"Street"`
	StreetNumber string `json:"StreetNumber"`
	StreetName   string `json:"StreetName"`
	UnitType     string `json:"UnitType"`
	UnitNumber   string `json:"UnitNumber"`
	City         string `json:"City"`
	Region       string `json:"Region"`
	PostalCode   string `json:"PostalCode"`
}

// Store represents a Domino's pizza store
type Store struct {
	StoreID             string  `json:"StoreID"`
	Phone               string  `json:"phone"`
	StreetName          string  `json:"StreetName'`
	City                string  `json:"City"`
	Region              string  `json:"Region"`
	PostalCode          string  `json:"PostalCode"`
	AddressDescription  string  `json:"AddressDescription"`
	IsDeliveryStore     bool    `json:"IsDeliveryStore"`
	MinDistance         float32 `json:"MinDistance"`
	MaxDistance         float32 `json:"MaxDistance"`
	HolidaysDescription string  `json:"HolidaysDescription"`
	Hours               struct {
		Sun struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Sun"`
		Mon struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Mon"`
		Tue struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Tue"`
		Wed struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Wed"`
		Thu struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Thu"`
		Fri struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Fri"`
		Sat struct {
			OpenTime  string `json:"OpenTime"`
			CloseTime string `json:"CloseTime"`
		} `json:"Sat"`
	}
	HoursDescription        string `json:"HoursDescription"`
	TimeZoneMinutes         int    `json:"TimeZoneMinutes"`
	ServiceHoursDescription struct {
		Carryout string `json:"Carryout"`
		Delivery string `json:"Delivery"`
	} `json:"ServiceHoursDescription"`
	IsOpen                          bool     `json:"IsOpen"`
	IsAffectedByDaylightSavingsTime bool     `json:"IsAffectedByDaylightSavingsTime"`
	IsOnlineCapable                 bool     `json:"IsOnlineCapable"`
	IsOnlineNow                     bool     `json:"IsOnlineNow"`
	IsNEONow                        bool     `json:"IsNEONow"`
	CashLimit                       int      `json:"CashLimit"`
	StoreAsOfTime                   string   `json:"StoreAsOfTime"`
	AsOfTime                        string   `json:"AsOfTime"`
	BusinessDate                    string   `json:"BusinessDate"`
	PreferredLanguage               string   `json:"PreferredLanguage"`
	PreferredCurrency               string   `json:"PreferredCurrency"`
	TimeZoneCode                    string   `json:"TimeZoneCode"`
	AcceptablePaymentTypes          []string `json:"AcceptablePaymentTypes"`
	AcceptableCreditCards           []string `json:"AcceptableCreditCards"`
	MinimumDeliveryOrderAmount      float32  `json:"MinimumDeliveryOrderAmount"`
	// LocationInfo unsure of value
	// LanguageLocationInfo
	AllowDeliveryOrders               bool `json:"AllowDeliveryOrders"`
	AllowCarryoutOrders               bool `json:"AllowCarryoutOrders"`
	ServiceMethodEstimatedWaitMinutes struct {
		Delivery struct {
			Min int `json:"Min"`
			Max int `json:"Max"`
		} `json:"Delivery"`
		Carryout struct {
			Min int `json:"Min"`
			Max int `json:"Max"`
		} `json:"Carryout"`
	} `json:"ServiceMethodEstimatedWaitMinutes"`
	StoreCoordinates struct {
		StoreLatitude  string `json:"StoreLatitude"`
		StoreLongitude string `json:"StoreLongitude"`
	} `json:"StoreCoordinates"`
	AllowPickupWindowOrders bool `json:"AllowPickupWindowOrders"`
	ServiceIsOpen           struct {
		Carryout bool `json:"Carryout"`
		Delivery bool `json:"Delivery"`
	} `json:"ServiceIsOpen"`
}

type CustomerOrder struct {
}

type Order struct {
	Address               Address   `json:"Address"`
	Coupons               []Coupon  `json:"Coupons"`
	CustomerID            string    `json:"CustomerID"`
	Email                 string    `json:"Email"`
	Extension             string    `json:"Extension"`
	FirstName             string    `json:"FirstName"`
	LastName              string    `json:"LastName"`
	OrderChannel          string    `json:"OrderChannel"`
	OrderID               string    `json:"OrderID"`
	OrderMethod           string    `json:"OrderMethod"`
	Payments              []string  `json:"Payments"`
	Phone                 string    `json:"Phone"`
	PhonePrefix           string    `json:"PhonePrefix"`
	ServiceMethod         string    `json:"ServiceMethod"`
	SourceOrganizationURI string    `json:"SourceOrganizationURI"`
	StoreID               string    `json:"StoreID"`
	Market                string    `json:"Market"`
	Current               string    `json:"Currency"`
	Products              []Product `json:"Products"`
}

type Coupon struct {
}

type Product struct {
}
