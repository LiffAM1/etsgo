package etsy

type User struct {
	CreationTsz  int64 `json:"creation_tsz"`
	FeedbackInfo struct {
		Count int64       `json:"count"`
		Score interface{} `json:"score"`
	} `json:"feedback_info"`
	LoginName        string      `json:"login_name"`
	ReferredByUserID interface{} `json:"referred_by_user_id"`
	UserID           int64       `json:"user_id"`
}
