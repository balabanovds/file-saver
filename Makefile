migration-create:
	goose -dir migrations sqlite3 ./db/db.sqlite create init sql

migrate-up:
	goose -dir migrations sqlite3 ./db/db.sqlite up

migrate-down:
	goose -dir migrations sqlite3 ./db/db.sqlite down
