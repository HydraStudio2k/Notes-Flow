module sqlite

go 1.22.2

replace db_models => ../../models

require (
	db_models v0.0.0-00010101000000-000000000000 // indirect
	github.com/mattn/go-sqlite3 v1.14.27 // indirect
)
