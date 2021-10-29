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
