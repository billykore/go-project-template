CREATE TABLE greets
(
    id         integer primary key autoincrement,
    name       string    not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp null,
    deleted_at timestamp null
);
