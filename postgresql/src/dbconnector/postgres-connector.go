package dbconnector

import (
	sql "database/sql"
	"encoding/json"
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

func UpdateSubArticle(id string, subArticleJson []byte) (int, error) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	sql := "UPDATE articles i SET subarticles = i2.subarticles :: jsonb || $2 :: jsonb FROM( SELECT id, array_to_json(array_agg(elem)) AS subarticles FROM articles i2, json_array_elements(i2.subarticles :: json) elem WHERE elem ->> 'id' <> $1 GROUP BY 1) i2 WHERE i2.id = i.id AND i.id = 1 AND json_array_length(i2.subarticles) < json_array_length(i.subarticles :: json) RETURNING $1"

	var subArticleId int
	err = db.QueryRow(sql, id, subArticleJson).Scan(&subArticleId)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return subArticleId, err
}

func DeleteSubArticle(id string) (int, error) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	sql := "UPDATE articles i SET subarticles = i2.subarticles FROM( SELECT id, array_to_json(array_agg(elem)) AS subarticles FROM articles i2, json_array_elements(i2.subarticles :: json) elem WHERE elem ->> 'id' <> $1 GROUP BY 1) i2 WHERE i2.id = i.id AND i.id = 1 AND json_array_length(i2.subarticles) < json_array_length(i.subarticles :: json) RETURNING $1"

	var subArticleId int
	err = db.QueryRow(sql, id).Scan(&subArticleId)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return subArticleId, err
}

func AddSubArticle(subArticleJson []byte) (int, error) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var subArticle SubArticle
	json.Unmarshal(subArticleJson, &subArticle)

	sql := "UPDATE articles SET subarticles = subarticles || $1 :: jsonb WHERE id = 1 RETURNING $2"

	var id int
	err = db.QueryRow(sql, subArticleJson, subArticle.ID).Scan(&id)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return id, err
}

func GetSubArticle(id string) (SubArticle, error) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var subArticle SubArticle
	sql := "SELECT items.id, items.published, items.authors, items.categories, items.tags FROM articles, jsonb_to_recordset(subarticles) AS items(id text, published bool, authors text, categories text, tags text) WHERE items.id = $1"

	err = db.QueryRow(sql, id).Scan(&subArticle.ID, &subArticle.Published, &subArticle.Authors, &subArticle.Categories, &subArticle.Tags)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return subArticle, err
}

func GetArticles() ([]Article, error) {
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

	return articles, err
}

func GetArticle(id int) (Article, error) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var article Article
	sql := "SELECT id, authors, subarticles FROM articles WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&article.ID, &article.Authors, &article.Subarticles)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return article, err
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
