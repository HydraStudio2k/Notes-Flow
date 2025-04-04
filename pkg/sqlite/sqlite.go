package sqlite

import (
	"log"

	models "db_models"
	// _ "github.com/mattn/go-sqlite3"
)

func Test_connection() {
	user := models.User{
		Id:       1,
		Name:     "user_test",
		Username: "@test",
		Password: "12345",
	}
	log.Println(user.Id, user.Name, user.Username, user.Password)
	log.Println("works")
}
