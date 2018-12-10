package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "jony:password@/test01")
	checkErr(err)
	defer db.Close()
	checkDBConn(db)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/read", read)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/del", del)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "at index")
	checkErr(err)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM amigos")
	checkErr(err)
	defer rows.Close()

	var s, pk, name string
	s = "\nRETRIEVED ROWS:\n"
	for rows.Next() {
		err = rows.Scan(&pk, &name)
		checkErr(err)
		s += pk + ":\t" + name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("CREATE TABLE customer (fname VARCHAR(20));")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)
	
	fmt.Fprint(w, "Table is created: ", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO customer (fname) values ('james');")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprint(w, "value james is inserted into customer table", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM customer;")
	checkErr(err)
	defer rows.Close()

	var fname string
	for rows.Next() {
		err = rows.Scan(&fname)
		checkErr(err)
		fmt.Fprint(w, "Retrived row from db customer: ", fname)
	}
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("DROP TABLE customer")
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprint(w, "table customer is dropped successfully", n)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET fname="Jony" WHERE fname="james";`)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprint(w, "in customer table value james is updated to Jony ", n)
}

func del(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE fname="Jony";`)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	checkErr(err)

	n, err := res.RowsAffected()
	checkErr(err)

	fmt.Fprint(w, "in customer table value james is updated to Jony ", n)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkDBConn(db *sql.DB) {
	if err = db.Ping(); err != nil {
		log.Fatal("database connection couldn't be established. error: ", err)
	}
}

