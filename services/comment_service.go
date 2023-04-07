package services

import (
	"github.com/naoymatsuda/myapi/apperrors"
	"github.com/naoymatsuda/myapi/models"
	"github.com/naoymatsuda/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
