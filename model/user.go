package model

import (
	"database/sql"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Verified bool
	Role     string
}

func initModelUser() {

}

func AddUser(user User) (ID int, err error) {
	query := "INSERT INTO users (id,username,password,email) VALUES (setserialval(),$1,$2,$3) RETURNING id;"
	err = PG.QueryRow(query, user.Username, user.Password, user.Email).Scan(&ID)
	return
}

func IsUserExistByName(name string) (is bool, err error) {
	query := "SELECT count(*) FROM users WHERE username = $1;"
	err = PG.QueryRow(query, name).Scan(&is)
	return
}

func IsUserExistByEmail(email string) (is bool, err error) {
	query := "SELECT count(*) FROM users WHERE email = $1;"
	err = PG.QueryRow(query, email).Scan(&is)
	return
}

func IsUserExistByID(ID int) (is bool, err error) {
	query := "SELECT count(*) FROM users WHERE id = $1;"
	err = PG.QueryRow(query, ID).Scan(&is)
	return
}

func GetUserByName(username string) (user User, is bool, err error) {
	query := "SELECT id,username,password,email,verified,role FROM users WHERE username = $1;"
	err = PG.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Verified, &user.Role)
	is = !(user.ID == 0)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func GetUserByID(ID int) (user User, is bool, err error) {
	query := "SELECT id,username,password,email,verified,role FROM users WHERE id = $1;"
	err = PG.QueryRow(query, ID).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Verified, &user.Role)
	is = !(user.ID == 0)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func GetAllUser() (users []User, err error) {
	query := "SELECT id,username,password,email,verified,role FROM users;"
	rows, err := PG.Query(query)
	var user User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Verified, &user.Role)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func IsUserAdmin(ID int) (is bool, err error) {
	query := "SELECT role FROM users WHERE id=$1;"
	var role string
	err = PG.QueryRow(query, ID).Scan(&role)
	is = role == "admin"
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func UpdateUser(ID int, user User) (err error) {
	var query string
	if user.Username != "" {
		if query != "" {
			query = query + ","
		}
		query = query + "username='" + user.Username + "'"
	}
	if user.Password != "" {
		if query != "" {
			query = query + ","
		}
		query = query + "password='" + user.Password + "'"
	}
	if user.Email != "" {
		if query != "" {
			query = query + ","
		}
		query = query + "email='" + user.Email + "'"
	}
	if user.Role != "" {
		if query != "" {
			query = query + ","
		}
		query = query + "role='" + user.Role + "'"
	}
	if query == "" {
		return nil
	}
	query = "UPDATE users SET " + query + " WHERE id = $1;"
	err = PG.QueryRow(query, ID).Scan()
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
