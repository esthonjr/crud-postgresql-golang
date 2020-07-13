// program to exemplify psql crud with golang
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// psql config
const (
	PGUSER     = "esthonjr"
	PGHOST     = "localhost"
	PGDATABASE = "testdb"
	PGPASSWORD = "estadmin"
	PGPORT     = 5432
)

func main() {
	opt := os.Args[1]

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	if opt == "create" {
		sqlStatement := `CREATE TABLE users (email varchar,firstName varchar,lastName varchar,age int)`
		_, err = db.Exec(sqlStatement)
		checkErr(err)
		fmt.Println("Table created")
	} else if opt == "read" {
		rows, err := db.Query("SELECT * FROM users WHERE age > 0")
		checkErr(err)

		for rows.Next() {
			var email string
			var firstName string
			var lastName string
			var age int
			err = rows.Scan(&email, &firstName, &lastName, &age)
			checkErr(err)
			fmt.Println("email | firstName | lastName | age ")
			fmt.Printf("%3v | %8v | %6v | %6v\n", email, firstName, lastName, age)
		}
	} else if opt == "update" {
		sqlStatement := `UPDATE users SET age = 72 WHERE email = 'test@test.com'`
		_, err = db.Exec(sqlStatement)
		checkErr(err)
		fmt.Println("Data updated")
	} else if opt == "delete" {
		sqlStatement := `DELETE FROM users WHERE email = 'test@test.com'`
		_, err = db.Exec(sqlStatement)
		checkErr(err)
		fmt.Println("Data deleted")
	} else if opt == "insert" {
		sqlStatement := `INSERT INTO users (email, firstName, lastName, age) VALUES ('test@test.com', 'test', 'test2', 33)`
		_, err = db.Exec(sqlStatement)
		checkErr(err)
		fmt.Println("Data inserted")
	} else if opt == "drop" {
		sqlStatement := `DROP TABLE users`
		_, err = db.Exec(sqlStatement)
		checkErr(err)
		fmt.Println("Table deleted")
	} else {
		fmt.Println(`Invalid option
	create # para criar a tabela
	read # para ler tabela
	update # para fazer update dos dados
	insert # para inserir dados
	delete # para deletar os dados
	drop # para deletar a tabela`)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
