package dbconnector
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "root"
    password = "root"
    dbname   = "postgres"
)

func initDB() {

}
 
func TestDB() {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
    db, err := sql.Open("postgres", psqlconn)
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