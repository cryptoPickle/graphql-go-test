package postgress

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host, user, dbName string, port int) string {
	return fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, user, dbName)
}


type User struct {
	ID int
	Name string
	Age string
	Friendly bool
}

func (db *Db) GetUsersByName(name string)[]User {
	stmt, err := db.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation err:", err)
	}

	rows, err := stmt.Query(name);

	if err != nil {
		fmt.Println("GetUserByName Query err:", err)
	}

	var r User

	users := []User{}

	for rows.Next() {
		err = rows.Scan(&r.ID, &r.Name, &r.Age, &r.Friendly)

		if err != nil {
			fmt.Println("GetUserByName Scan err:", err)
		}

		users = append(users, r)
	}

	return users
}