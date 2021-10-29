package model

import (
	"database/sql"
	"strconv"
)

func AddVerifyCode(code string, user User) (err error) {
	query := "INSERT INTO email (id,email,code) VALUES ($1,$2,$3);"
	err = PG.QueryRow(query, user.ID, user.Email, code).Scan()
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
	query := "UPDATE users SET verified=" + strconv.FormatBool(verified) + " WHERE id = $1;"
	err = PG.QueryRow(query, ID).Scan()
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
