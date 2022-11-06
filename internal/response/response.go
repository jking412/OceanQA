package response

type QuestionListResp struct {
	QuestionList []QuestionResp `json:"question_list"`
	Pagination   Pagination     `json:"pagination"`
}

type QuestionResp struct {
	Id            uint64 `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	ReadTime      uint64 `json:"read_time"`
	Title         string `json:"title"`
	OriginContent string `json:"origin_content"`
	ParseContent  string `json:"parse_content"`
	AnswerUrl     string `json:"answer_url"`
	Tags          []Tag  `json:"tags"`
	IsStared      bool   `json:"is_stared"`
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
