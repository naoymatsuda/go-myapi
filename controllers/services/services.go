package services

import (
	"github.com/naoymatsuda/myapi/models"
)

type ArticleServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
