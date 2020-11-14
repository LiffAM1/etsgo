package etsy

type Shop struct {
	AcceptsCustomRequests          bool        `json:"accepts_custom_requests"`
	Announcement                   interface{} `json:"announcement"`
	CreationTsz                    int64       `json:"creation_tsz"`
	CurrencyCode                   string      `json:"currency_code"`
	DigitalListingCount            int64       `json:"digital_listing_count"`
	DigitalSaleMessage             interface{} `json:"digital_sale_message"`
	HasOnboardedStructuredPolicies bool        `json:"has_onboarded_structured_policies"`
	HasUnstructuredPolicies        bool        `json:"has_unstructured_policies"`
	IconURLFullxfull               interface{} `json:"icon_url_fullxfull"`
	ImageURL760x100                interface{} `json:"image_url_760x100"`
	IncludeDisputeFormLink         bool        `json:"include_dispute_form_link"`
	IsCalculatedEligible           bool        `json:"is_calculated_eligible"`
	IsDirectCheckoutOnboarded      bool        `json:"is_direct_checkout_onboarded"`
	IsOptedInToBuyerPromise        bool        `json:"is_opted_in_to_buyer_promise"`
	IsShopUsBased                  bool        `json:"is_shop_us_based"`
	IsUsingStructuredPolicies      bool        `json:"is_using_structured_policies"`
	IsVacation                     bool        `json:"is_vacation"`
	Languages                      []string    `json:"languages"`
	LastUpdatedTsz                 int64       `json:"last_updated_tsz"`
	ListingActiveCount             int64       `json:"listing_active_count"`
	LoginName                      string      `json:"login_name"`
	NumFavorers                    int64       `json:"num_favorers"`
	PolicyAdditional               interface{} `json:"policy_additional"`
	PolicyHasPrivateReceiptInfo    bool        `json:"policy_has_private_receipt_info"`
	PolicyPayment                  interface{} `json:"policy_payment"`
	PolicyPrivacy                  interface{} `json:"policy_privacy"`
	PolicyRefunds                  interface{} `json:"policy_refunds"`
	PolicySellerInfo               interface{} `json:"policy_seller_info"`
	PolicyShipping                 interface{} `json:"policy_shipping"`
	PolicyUpdatedTsz               int64       `json:"policy_updated_tsz"`
	PolicyWelcome                  interface{} `json:"policy_welcome"`
	SaleMessage                    interface{} `json:"sale_message"`
	ShopID                         int64       `json:"shop_id"`
	ShopName                       string      `json:"shop_name"`
	Title                          interface{} `json:"title"`
	UpcomingLocalEventID           interface{} `json:"upcoming_local_event_id"`
	URL                            string      `json:"url"`
	UserID                         int64       `json:"user_id"`
	VacationAutoreply              interface{} `json:"vacation_autoreply"`
	VacationMessage                interface{} `json:"vacation_message"`
}
