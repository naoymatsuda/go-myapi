package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/naoymatsuda/myapi/services"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", "../sampleDB.sql")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)

	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}

