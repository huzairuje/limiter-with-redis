package article

import (
	"context"
	"fmt"
	"strings"

	"github.com/test_cache_CQRS/module/primitive"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateArticle(ctx context.Context, payload primitive.Article) (primitive.Article, error)
	CountArticle(ctx context.Context, param primitive.ParameterFindArticle) (int64, error)
	FindListArticle(ctx context.Context, param primitive.ParameterFindArticle) ([]primitive.Article, error)
	FindArticleByID(ctx context.Context, articleID int64) (primitive.Article, error)
	SetParamQueryToOrderByQuery(orderBy string) string
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateArticle(ctx context.Context, payload primitive.Article) (primitive.Article, error) {
	if err := r.db.WithContext(ctx).Table("articles").Create(&payload).Error; err != nil {
		return payload, err
	}
	return payload, nil
}

func (r *Repository) CountArticle(ctx context.Context, param primitive.ParameterFindArticle) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Table("articles")
	if param.Author != "" {
		query.Where(`"author" ILIKE ?`, "%"+param.Author+"%")
	}
	if param.Query != "" {
		query.Where(`"title" ILIKE ? or "body" ILIKE ?`, "%"+param.Query+"%", "%"+param.Query+"%")
	}
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) FindListArticle(ctx context.Context, param primitive.ParameterFindArticle) ([]primitive.Article, error) {
	var listData []primitive.Article
	query := r.db.WithContext(ctx).Table("articles")
	query.Where(`"deleted_at" is null`)
	if param.Author != "" {
		query.Where(`"author" ILIKE ?`, "%"+param.Author+"%")
	}
	if param.Query != "" {
		query.Where(`"title" ILIKE ? or "body" ILIKE ?`, "%"+param.Query+"%", "%"+param.Query+"%")
	}
	err := query.Offset(param.Offset).
		Limit(param.PageSize).
		Order(strings.Join([]string{param.SortBy, param.SortOrder}, " ")).
		Find(&listData).
		Error
	if err != nil {
		return nil, err
	}
	return listData, nil
}

func (r *Repository) SetParamQueryToOrderByQuery(orderBy string) string {
	var result string
	switch orderBy {
	case "id":
		result = fmt.Sprintf(`id`)
	case "author":
		result = fmt.Sprintf(`author`)
	case "title":
		result = fmt.Sprintf(`title`)
	case "body":
		result = fmt.Sprintf(`body`)
	case "created":
		result = fmt.Sprintf(`created_at`)
	default:
		result = fmt.Sprintf(`created_at`)
	}
	return result
}

func (r *Repository) FindArticleByID(ctx context.Context, articleID int64) (primitive.Article, error) {
	var data primitive.Article
	err := r.db.WithContext(ctx).
		Table("articles").
		Where(`"deleted_at" is null and id = ?`, articleID).
		First(&data).
		Error
	if err != nil {
		return primitive.Article{}, err
	}
	return data, nil
}
