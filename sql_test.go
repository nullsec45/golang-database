package main

import (
	"context"
	"database/sql"
	"fmt"
	. "golang-database/database"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO customer(id, name) VALUES ('01','Fajar')"
	_, err := db.ExecContext(ctx, querySql)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		var err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("Birth Date :", birthDate)
		fmt.Println("Married :", married)
		fmt.Println("Created At", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "nullsec45';#"
	password := "rahasiaanget"
	querySql := "SELECT username FROM users WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Selamat datang", username)
	} else {
		fmt.Println("Gagal Login!")
	}
}

func TestSqlSecure(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "nullsec45';#"
	password := "rahasiaanget"
	querySql := "SELECT username FROM users WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, querySql, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Selamat datang", username)
	} else {
		fmt.Println("Gagal Login!")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "nullsec45';DROP TABLE users#"
	password := "rahasiaanget"

	querySql := "INSERT INTO users(username,password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, querySql, username, password)

	if err != nil {
		panic(err)
	}
	fmt.Println("Sukes menambahkan user baru")
}

func TestExecSqlLastInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "ramafajar805@gmail.com"
	comment := "mantap"

	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
	fmt.Println("Last insert id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments (email, comment) VALUES (?,?)"
	stmt, err := db.PrepareContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "rama" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}
		LastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", LastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments (email, comment) VALUES (?,?)"
	for i := 0; i < 10; i++ {
		email := "rama" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)

		if err != nil {
			panic(err)
		}
		LastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", LastInsertId)
	}
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
