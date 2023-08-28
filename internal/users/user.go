package users

import (
	"database/sql"
	"example/graph/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Person struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

func (user *Person) Authenticate() bool {
	db, err := model.InitDB("root:lokeshpathrabe@tcp(localhost:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("select password from persons WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	db, err := model.InitDB("root:lokeshpathrabe@tcp(localhost:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("select id from persons WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}