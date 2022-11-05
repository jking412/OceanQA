package response

type QuestionResponse struct {
	QuestionList []Question `json:"question_list"`
	Pagination   Pagination `json:"pagination"`
}

type Question struct {
	Id        uint64   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	IsStared  bool     `json:"is_stared"`
}

type Pagination struct {
	PageNum  uint64 `json:"page_num"`
	PageSize uint64 `json:"page_size"`
	Total    uint64 `json:"total"`
}

type Tag struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
