package types

type StrollReplay struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	BV        string `json:"bv"`
	TargetUrl string `json:"target_url"`
	CreatedAt uint   `json:"created_at"`
}
