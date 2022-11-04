package request

type AllQuestionRequest struct {
	CurrentNum uint64 `json:"current_num"`
	PageSize   uint64 `json:"page_size"`
}
