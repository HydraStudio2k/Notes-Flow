module sqlite

go 1.22.2

replace db_models => ../../models

require (
	db_models v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.27 // indirect
	golang.org/x/text v0.14.0 // indirect
)
