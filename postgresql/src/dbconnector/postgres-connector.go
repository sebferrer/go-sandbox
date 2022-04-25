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

func openConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", psqlconn)
}

func GetArticles() {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var myArticle Article
	sql := "SELECT id, authors, subarticles FROM articles"

	err = db.QueryRow(sql).Scan(&myArticle.ID, &myArticle.Authors, &myArticle.Subarticles)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("%s\n", myArticle)
}

func GetArticle(id int) {
	db, err := openConnection()
	CheckError(err)
	defer db.Close()

	var myArticle Article
	sql := "SELECT id, authors, subarticles FROM articles WHERE id = $1"

	err = db.QueryRow(sql, "1").Scan(&myArticle.ID, &myArticle.Authors, &myArticle.Subarticles)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("%s\n", myArticle)
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
