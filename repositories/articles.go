package repositories

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoymatsuda/myapi/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice) values (?, ?, ?, ?);
	`

	var newArticle models.Article

	newArticle.Title, newArticle.Contents, newArticle.UserName, newArticle.NiceNum = article.Title, article.Contents, article.UserName, article.NiceNum
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName, article.NiceNum)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {

	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return nil, err
		}
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow("select * from articles where article_id = ?;", 1)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlStr = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `
		update articles
		set nice = ?
		where article_id = ?;
	`
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if tx.Commit(); err != nil {
		return err
	}

	return nil
}
