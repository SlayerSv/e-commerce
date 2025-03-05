drop table if exists users cascade;
create table users (
    id serial primary key,
    name varchar(30) unique not null,
    password text not null
);