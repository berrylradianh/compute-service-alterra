package blogs

type BlogResponse struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
}

func (BlogResponse) TableName() string {
	return "blogs"
}
