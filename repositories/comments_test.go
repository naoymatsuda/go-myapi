package repositories_test

import (
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoymatsuda/myapi/models"
	"github.com/naoymatsuda/myapi/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != expectedNum {
		t.Errorf("get %d but want %d\n", len(got), expectedNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
		CreatedAt: time.Now(),
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)

	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	// t.Cleanup(func() {
	// 	const sqlStr = `DELETE FROM articles WHERE title = ? and contents = ? and user_name = ?`
	// 	testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	// })
}
