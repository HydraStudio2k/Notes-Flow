package sqlite

import (
	"log"

	models "db_models"

	"github.com/google/uuid"
)

func Test_connection() {
	user := models.User{
		Id:       1,
		Name:     "user_test",
		Username: "@test",
		Password: "12345",
		Uuid:     uuid.NewString(),
	}
	note := models.Note{
		Id:     1,
		Author: "Test_User",
		Status: "Public/Private",
		Note:   "#1",
		Uuid:   uuid.NewString(),
	}
	log.Println("[Test_user info]", user.Id, user.Name, user.Username, user.Password, user.Uuid)
	log.Println("[Test_note info]", note.Id, note.Author, note.Status, note.Note, note.Uuid)
	log.Println("works")
}
