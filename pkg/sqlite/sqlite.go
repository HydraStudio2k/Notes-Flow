package sqlite

import (
	"log"
	"path/filepath"

	models "db_models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
|--------------------|
| DATABASE FUNCTIONS |
|--------------------|
*/

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

/*
|---------------- |
| USERS FUNCTIONS |
|---------------- |
*/

// +
func CreateUser(name string, username string, password string) bool { // Adds a new user to the sqlite database. Returns true if the user was added successfully.
	// To add, you need to come up with a unique username; if the username already exists, adding to the database is impossible.

	var users []models.User

	db := DBInit()

	db.Where("username = ?", username).Find(&users)

	if len(users) == 0 { // If there are no users with such a username, addiing a new user.
		user_uuid := uuid.NewString() // Create a unique user id.
		db.Create(&models.User{
			Name:     name,
			Username: username,
			Password: password,
			Uuid:     user_uuid,
		})
		DBClose(db)
		return true
	} else {
		DBClose(db)
		return false
	}
}

func UserExistence(username string, password string) bool { // Returns true if such user exists.
	var user models.User

	db := DBInit()

	err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
	DBClose(db)

	if err != nil {
		return false
	} else {
		return true
	}
}

/*
|-----------------|
| NOTES FUNCTIONS |
|-----------------|
*/

func CreateNote(author string, status string, note string) bool { // Adds a new note to the sqlite database. Returns true if the note was added successfully.
	// The add only exists if a user with the specified name already exists in the database.

	var notes []models.Note

	db := DBInit()

	db.Where("username = ?", author).Find(&notes)

	if len(notes) == 0 { // If there are no users with that name, the note will not be added.
		DBClose(db)

		return false
	} else {
		note_uuid := uuid.NewString() // Create a unique note id.

		db.Create(&models.Note{
			Author: author,
			Status: status,
			Note:   note,
			Uuid:   note_uuid,
		})

		DBClose(db)

		return true
	}
}

func GetNoteByUUID(uuid string) (models.Note, error) { // Returns a note object if the specified identifier actually exists, otherwise if nothing was found an empty object and an error will be returned.
	var note models.Note
	db := DBInit()

	err := db.Where("uuid = ?", uuid).First(&note).Error
	DBClose(db)

	if err != nil {
		return note, err
	} else {
		return note, nil
	}
}

func GetPublicNotesByAuthor(author string) ([]models.Note, error) { // Returns a slice of note objects associated with a user with public status, if it actually exists, otherwise an empty slice and an error are returned.
	var notes []models.Note
	db := DBInit()

	err := db.Where("author = ? AND status = ?", author, "public").Find(&notes).Error
	DBClose(db)

	if err != nil {
		return notes, err
	} else {
		return notes, nil
	}
}

func GetPrivateNotesByAuthor(author string) ([]models.Note, error) { // Returns a slice of private notes belonging to a specific author. On error, returns an empty slice and an error.
	var notes []models.Note
	db := DBInit()

	err := db.Where("author = ? AND status = ?", author, "private").Find(&notes).Error
	DBClose(db)

	if err != nil {
		return notes, err
	} else {
		return notes, nil
	}
}

func GetAllPublicNotes() ([]models.Note, error) { // Returns a slice of note objects with the status "public", otherwise returns an empty slice and an error.
	var notes []models.Note
	db := DBInit()

	err := db.Where("status = ?", "public").Find(&notes).Error
	DBClose(db)

	if err != nil {
		return notes, err
	} else {
		return notes, nil
	}
}

func DeleteNoteByUUID(uuid string) error { // Deletes a note by the specified uuid.
	db := DBInit()

	err := db.Where("uuid = ?", uuid).Delete(&models.Note{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
