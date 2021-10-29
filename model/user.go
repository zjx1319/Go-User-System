package model

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
	sql := "INSERT INTO users (id,username,password,email) VALUES (setserialval(),$1,$2,$3) RETURNING id;"
	err = PG.QueryRow(sql, user.Username, user.Password, user.Email).Scan(&ID)
	return
}

func IsUserExistByName(name string) (is bool, err error) {
	sql := "SELECT count(*) FROM users WHERE username = $1"
	err = PG.QueryRow(sql, name).Scan(&is)
	return
}
