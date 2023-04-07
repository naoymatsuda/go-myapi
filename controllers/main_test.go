package controllers_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoymatsuda/myapi/controllers"
	"github.com/naoymatsuda/myapi/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
