module main

go 1.22.2

require github.com/gorilla/mux v1.8.1

require main/handlers v0.0.0-00010101000000-000000000000

require (
	db_models v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.27 // indirect
	golang.org/x/text v0.14.0 // indirect
	gorm.io/driver/sqlite v1.5.7 // indirect
	gorm.io/gorm v1.25.12 // indirect
	handlers/sqlite v0.0.0-00010101000000-000000000000 // indirect
)

replace main/handlers => ../../pkg/handlers // allow access to handlers

replace handlers/sqlite => ../../pkg/sqlite // allow access to sqlite

replace db_models => ../../models
