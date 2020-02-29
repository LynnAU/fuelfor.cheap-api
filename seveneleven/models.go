package seveneleven

// some fields have been commented out because they are not relevant to this application

// LoginPayload ...
type LoginPayload struct {
	Email               string `json:"Email"`
	Password            string `json:"Password"`
	DeviceName          string `json:"DeviceName"`
	DeviceOSNameVersion string `json:"DeviceOsNameVersion"`
}

// LoginResponse ...
type LoginResponse struct {
	DeviceSecretToken string `json:"DeviceSecretToken"`
	AccountID         string `json:"AccountId"`
	FirstName         string `json:"FirstName"`
	Email             string `json:"Email"`
	// TouchBaseURL      string                    `json:"TouchBaseURL"`
	DigitalCard *LoginResponseDigitalCard `json:"DigitalCard"`
}

// LoginResponseDigitalCard ...
type LoginResponseDigitalCard struct {
	// CardNumber int     `json:"CardNumber"`
	Balance float64 `json:"Balance"`
	// CardSecret string  `json:"CardSecret"`
	// Name       string  `json:"Name"`
	// Expiry     float64 `json:"Expiry"`
}

// LockSessionResponse ...
type LockSessionResponse struct {
	SessionExpiry int `json:"SessionExpiry"`
	// StoreUpdates  struct {
	// 	AsOfDate int           `json:"AsOfDate"`
	// 	Diffs    []interface{} `json:"Diffs"`
	// } `json:"StoreUpdates"`
	CheapestFuelTypeStores []struct {
		StoreNumber string `json:"StoreNumber"`
		FuelPrices  []struct {
			Ean       int     `json:"Ean"`
			Price     float64 `json:"Price"`
			PriceDate int     `json:"PriceDate"`
		} `json:"FuelPrices"`
	} `json:"CheapestFuelTypeStores"`
	Balance float64 `json:"Balance"`
}
