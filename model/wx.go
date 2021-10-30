package model

import "database/sql"

func BindWX(ID int, wxName string, openid string) (err error) {
	query := "INSERT INTO wx (id,wxname,openid) VALUES ($1,$2,$3) ON CONFLICT(id) DO UPDATE SET wxname=$2,openid=$3;"
	err = PG.QueryRow(query, ID, wxName, openid).Scan()
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func GetUserByWX(openid string) (user User, is bool, err error) {
	query := "SELECT id,username,password,email,verified,role FROM users JOIN wx USING(id) WHERE openid=$1;"
	err = PG.QueryRow(query, openid).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Verified, &user.Role)
	is = !(user.ID == 0)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func GetWXName(ID int) (WXName string, is bool, err error) {
	query := "SELECT wxname FROM wx WHERE id=$1;"
	err = PG.QueryRow(query, ID).Scan(&WXName)
	is = !(WXName == "")
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
