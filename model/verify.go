package model

import (
	"database/sql"
)

func AddVerifyCode(code string, ID int, Email string) (err error) {
	query := "INSERT INTO email (id,email,code) VALUES ($1,$2,$3);"
	err = PG.QueryRow(query, ID, Email, code).Scan()
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func GetVerifyCode(ID int) (code string, is bool, err error) {
	query := "SELECT code FROM email WHERE id = $1"
	err = PG.QueryRow(query, ID).Scan(&code)
	is = !(code == "")
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func DeleteVerifyCode(code string) (err error) {
	query := "DELETE FROM email WHERE code = $1"
	err = PG.QueryRow(query, code).Scan(&code)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func UpdateUserVerified(ID int, verified bool) (err error) {
	query := "UPDATE users SET verified=$1 WHERE id = $2;"
	err = PG.QueryRow(query, verified, ID).Scan()
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
