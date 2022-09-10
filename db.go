package main

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	
)
func main() {
	connStr := "postgresql://postgres:postgres@localhost/news?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()
	res, err := db.Query(`INSERT INTO "Newskg" (id,news_theme,news_paragraph,news_link,news_site) VALUES (3,'dgsfsdfsdffsdf sdf sdf','sdfsdf','sdfsdf','sdfsdfdsf')`)
    if err != nil{    
		panic(err)
    } 
	fmt.Println(res)
}