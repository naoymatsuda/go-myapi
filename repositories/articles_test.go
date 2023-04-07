package repositories_test

import (
	// "testing"

	// _ "github.com/mattn/go-sqlite3"
	// "github.com/naoymatsuda/myapi/models"
	// "github.com/naoymatsuda/myapi/repositories"
)

// func TestSelectArticleDetail(t *testing.T) {

// 	tests := []struct {
// 		testTitle string
// 		expected  models.Article
// 	}{
// 		{
// 			testTitle: "subTest1",
// 			expected: models.Article{
// 				ID:       1,
// 				Title:    "2nd",
// 				Contents: "Second blog post",
// 				UserName: "saki",
// 				NiceNum:  11,
// 			},
// 		}, {
// 			testTitle: "subTest2",
// 			expected: models.Article{
// 				ID:       2,
// 				Title:    "1st",
// 				Contents: "1st blog post",
// 				UserName: "saki",
// 				NiceNum:  2,
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.testTitle, func(t *testing.T) {
// 			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			if got.ID != test.expected.ID {
// 				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
// 			}
// 			if got.Title != test.expected.Title {
// 				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
// 			}
// 			if got.Contents != test.expected.Contents {
// 				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
// 			}
// 			if got.UserName != test.expected.UserName {
// 				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
// 			}
// 			if got.NiceNum != test.expected.NiceNum {
// 				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
// 			}
// 		})
// 	}
// }

// func TestSelectArticleList(t *testing.T) {

// 	expectedNum := 3
// 	got, err := repositories.SelectArticleList(testDB, 1)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if num := len(got); num != expectedNum {
// 		t.Errorf("want %d but got %d articles\n", expectedNum, num)
// 	}

// }

// func TestInsertArticle(t *testing.T) {
// 	article := models.Article{
// 		Title:    "Insert Test",
// 		Contents: "test",
// 		UserName: "naoyuki",
// 		NiceNum:  0,
// 	}

// 	expectedArticleNum := 10
// 	newArticle, err := repositories.InsertArticle(testDB, article)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if newArticle.ID != expectedArticleNum {
// 		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
// 	}

// 	t.Cleanup(func() {
// 		const sqlStr = `DELETE FROM articles WHERE title = ? and contents = ? and user_name = ?`
// 		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
// 	})
// }

// func TestUpdateNiceNum(t *testing.T) {
// 	articleID := 1
// 	before, err := repositories.SelectArticleDetail(testDB, articleID)
// 	if err != nil {
// 		t.Fatal("fail to get before data")
// 	}

// 	err = repositories.UpdateNiceNum(testDB, articleID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	after, err := repositories.SelectArticleDetail(testDB, articleID)
// 	if err != nil {
// 		t.Fatal("fail to get after data")
// 	}

// 	if after.NiceNum - before.NiceNum != 1 {
// 		t.Error("fail to update nice num")

// 	}
// }
