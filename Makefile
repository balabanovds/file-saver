migration-create:
	goose -dir migrations sqlite3 ./db/db.sqlite create init sql

migrate-up:
	goose -dir migrations sqlite3 ./db/db.sqlite up

migrate-down:
	goose -dir migrations sqlite3 ./db/db.sqlite down

clean-db: migrate-down migrate-up

build: clean
	GOOS=darwin GOARCH=amd64 go build -o ./build/file-saver_darwin_amd64 ./cmd

build-linux: clean
	GOOS=linux GOARCH=amd64 go build -o ./build/file-saver_linux_amd64 ./cmd

clean:
	rm -rf ./build

pack: clean-db build-linux
	mv ./build/file-saver_linux_amd64 .
	tar czf app.tgz file-saver_linux_amd64 ./db
	mv app.tgz ./build
	rm -rf file-saver_linux_amd64
