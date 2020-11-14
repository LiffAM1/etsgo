package etsy

type Listing struct {
	IsVintage bool          `json:"is_vintage"`
	ListingID int64         `json:"listing_id"`
	Sku       []interface{} `json:"sku"`
	State     string        `json:"state"`
}
