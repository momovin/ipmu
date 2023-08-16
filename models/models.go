package models

type locationStruct struct {
	Capital         string `json:"capital"`
	LanguagesCode   string `json:"languages_code"`
	LanguagesName   string `json:"languages_name"`
	LanguagesNative string `json:"languages_native"`
	CountryFlag     string `json:"country_flag"`
	CountryUnicode  string `json:"country_unicode"`
	CallingCode     string `json:"calling_code"`
	IsEU            bool   `json:"is_eu"`
}

type timeZoneStruct struct {
	ID          string `json:"id"`
	CurrentTime string `json:"current_time"`
	Code        string `json:"code"`
	UtcOffset   string `json:"utc_offset"`
}

type currencyStruct struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type connectionStruct struct {
	ASN      string `json:"asn"`
	ISP      string `json:"isp"`
	Org      string `json:"org"`
	Registry string `json:"registry"`
}

type securityStruct struct {
	VPN     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Relay   bool `json:"relay"`
	Hosting bool `json:"hosting"`
}

type APIStruct struct {
	IP                string  `json:"ip"`
	Type              string  `json:"type"`
	ContinentCode     string  `json:"continent_code"`
	ContinentName     string  `json:"continent_name"`
	CountryCode       string  `json:"country_code"`
	CountryName       string  `json:"country_name"`
	RegionCode        string  `json:"region_code"`
	RegionName        string  `json:"region_name"`
	City              string  `json:"city"`
	Zip               string  `json:"zip"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	*locationStruct   `json:"location"`
	*timeZoneStruct   `json:"time_zone"`
	*currencyStruct   `json:"currency"`
	*connectionStruct `json:"connection"`
	*securityStruct   `json:"security"`
	UserAgent         string `json:"user_agent"`
}
