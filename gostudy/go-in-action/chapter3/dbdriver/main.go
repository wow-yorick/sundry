package main

import (
	"database/sql"
	"fmt"
	_ "sundry/gostudy/go-in-action/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
	fmt.Println("xxx")
}
