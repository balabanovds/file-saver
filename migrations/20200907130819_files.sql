-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table files (
   id integer primary key autoincrement unique ,
   os_name text,
   app_name text,
   size integer,
   os_created_at timestamp,
   created_at timestamp default current_timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table files;
