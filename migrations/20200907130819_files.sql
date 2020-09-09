-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table files (
   id integer primary key autoincrement unique ,
   os_name text,
   app_name text,
   inode integer
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table files;
