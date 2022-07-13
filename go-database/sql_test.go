package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('ihab', 'Ihab')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}

	fmt.Println("Success insert new customer")

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating int64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===================")
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birthDate: ", birthDate.Time)

		}
		fmt.Println("married: ", married)

		fmt.Println("createdAt: ", createdAt)
	}

	fmt.Println("Success insert new customer")

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "y"
	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username: ", username)
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login gagal")

	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "y"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username: ", username)
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login gagal")

	}

}

func TestExecSqlSafeParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "imam'; DROP TABLE user; #"
	password := "imam"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}
