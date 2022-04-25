package dbconnector

import (
	sql "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

type Article struct {
	ID          string
	Authors     string
	Subarticles string
}

type SubArticle struct {
	ID         string
	Published  string
	Authors    string
	Categories string
	Tags       string
}

func openConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", psqlconn)
}

func GetSubArticle(id string) SubArticle {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var subArticle SubArticle
	sql := "select items.id, items.published, items.authors, items.categories, items.tags from articles, jsonb_to_recordset(subarticles) as items(id text, published bool, authors text, categories text, tags text) where items.id = $1"

	err = db.QueryRow(sql, id).Scan(&subArticle.ID, &subArticle.Published, &subArticle.Authors, &subArticle.Categories, &subArticle.Tags)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// fmt.Printf("%s\n", subArticle)
	return subArticle
}

func GetArticles() []Article {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var articles []Article

	sql := "SELECT id, authors, subarticles FROM articles"

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	for rows.Next() {
		var article Article
		rows.Scan(&article.ID, &article.Authors, &article.Subarticles)
		articles = append(articles, article)
	}

	// fmt.Printf("%s\n", articles)
	return articles
}

func GetArticle(id int) Article {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var article Article
	sql := "SELECT id, authors, subarticles FROM articles WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&article.ID, &article.Authors, &article.Subarticles)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// fmt.Printf("%s\n", article)
	return article
}

func TestDB() {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
