package primitive

type ArticleReq struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}
