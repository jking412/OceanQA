package request

type AllQuestionRequest struct {
	CurrentNum uint64 `json:"current_num"`
	PageSize   uint64 `json:"page_size"`
}

type CreateQuestionRequest struct {
	Title         string `json:"title"`
	AnswerUrl     string `json:"answer_url"`
	OriginContent string `json:"origin_content"`
	ParseContent  string `json:"parse_content"`
	Tags          []Tag  `json:"tags"`
}

type Tag struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
