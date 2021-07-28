package types

type PaginationList struct {
	List    interface{} `json:"list"`
	NextKey interface{} `json:"next_key"`
}

type StrollReply struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	BV        string `json:"bv"`
	PlayUrl   string `json:"play_url"`
	TargetUrl string `json:"target_url"`
	CreatedAt uint   `json:"created_at"`
}

type NextGroupReq struct {
	NextKey uint `form:"next_key,default=0" binding:"omitempty,numeric,gte=0"`
	Size    uint `form:"size,default=50" binding:"omitempty,numeric,gt=0,lt=100"`
}

type NextGroupReply struct {
	Title     string `json:"title"`
	Subtitled string `json:"subtitled"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	TargetUrl string `json:"target_url"`
	Timestamp uint   `json:"timestamp"`
}
