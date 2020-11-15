package etsy

type Country struct {
	CountryID            int64   `json:"country_id"`
	IsoCountryCode       string  `json:"iso_country_code"`
	Lat                  float64 `json:"lat"`
	Lon                  float64 `json:"lon"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	WorldBankCountryCode string  `json:"world_bank_country_code"`
}
