package article

import (
	"database/sql"
	"html/template"

	m "populi/model"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func (a *m.Article) save() error {
	// write to database here
	return nil
}

func loadArticle(id string) (*m.Article, error) {
	return nil, nil
}
