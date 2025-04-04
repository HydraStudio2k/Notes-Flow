module handlers

go 1.22.2

replace handlers/sqlite => ../sqlite // allow access to sqlite

replace db_models => ../../models

require handlers/sqlite v0.0.0-00010101000000-000000000000

require db_models v0.0.0-00010101000000-000000000000 // indirect
