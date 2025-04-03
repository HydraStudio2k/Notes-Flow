module main

go 1.22.2

require github.com/gorilla/mux v1.8.1

require main/handlers v0.0.0-00010101000000-000000000000

require handlers/sqlite v0.0.0-00010101000000-000000000000 // indirect

replace main/handlers => ../../pkg/handlers // allow access to handlers

replace handlers/sqlite => ../../pkg/sqlite // allow access to sqlite
