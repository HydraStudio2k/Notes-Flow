package sqlite

import (
	"log"
	"path/filepath"

	models "db_models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_connection() { // Test function for checks, in the future it will be removed.
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

	db := DBInit()
	DBClose(db)

	log.Println("[Test_user info]", user.Id, user.Name, user.Username, user.Password, user.Uuid)
	log.Println("[Test_note info]", note.Id, note.Author, note.Status, note.Note, note.Uuid)
	log.Println("works")
}

func DBInit() *gorm.DB { // Initializes a connection to a sqlite database and returns a database object for further manipulation. Otherwise, if an error occurs, returns nil and logs the error.
	db, err := gorm.Open(sqlite.Open(filepath.Join("..", "..", "pkg", "sqlite", "db", "sqlite.db")), &gorm.Config{})
	if err != nil {
		log.Printf("Error while opening sqlite database in sqlite.go: %e\n", err)
		return nil
	}

	db.AutoMigrate(&models.User{}, &models.Note{})
	return db
}

func DBClose(db *gorm.DB) error { // Closes the connection to the sqlite database, otherwise returns an error and logs it.
	sqlite_db, err := db.DB()
	if err != nil {
		log.Printf("Error while getting sqlite database in sqlite.go: %e\n", err)
		return err
	}

	err = sqlite_db.Close()
	if err != nil {
		log.Printf("Error while closing sqlite database in sqlite.go %e\n", err)
		return err
	}
	return nil
}
