package model

func AddVerifyCode(code string, user User) (err error) {
	sql := "INSERT INTO email (id,email,code) VALUES ($1,$2,$3);"
	err = PG.QueryRow(sql, user.ID, user.Email, code).Scan()
	return
}
