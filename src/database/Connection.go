package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func testConnection() *sql.DB{
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func createUserTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER NOT NULL AUTO_INCREMENT,userName varchar(255),password varchar(20),urlNum int,PRIMARY KEY (id))")
	if err0 != nil {
		log.Fatal(err0.Error())
	} else {
		_, err1 := st.Exec()
		if err1 != nil {
			log.Fatal(err1.Error())
		}
	}
}


func InsertNewUser(db *sql.DB, userName string, password string) {

	st, err0 := db.Prepare("insert into users (userName, password, urlNum) values (?,?,?)")
	if err0 != nil {
		panic(err0.Error())
	}

	_, err1 := st.Exec(userName, password, 0)
	if err1 != nil {
		panic(err1.Error())
	}

}

func createUrlsTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER NOT NULL AUTO_INCREMENT,url varchar(255),userName varchar(255),HealthCheck int,respOkTime int,respWarTime int,respCritTime int,PRIMARY KEY (id))")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}


func insertNewUrl(db *sql.DB) {
	st,err0 := db.Prepare("insert into urls()")
}


