package primitive

import "time"

type Article struct {
	ID        int64     `gorm:"column:id"`
	Author    string    `gorm:"column:author"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type ParameterFindArticle struct {
	Query     string
	Author    string
	PageSize  int
	Offset    int
	SortBy    string
	SortOrder string
}

type ParameterArticleHandler struct {
	Query  string
	Author string
}
