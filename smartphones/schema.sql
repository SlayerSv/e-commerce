drop table if exists smartphones cascade;
create table smartphones (
    id serial primary key,
    model varchar(20) not null,
    producer varchar(20) not null,
    color varchar(20) not null,
    screen_size numeric(3,2),
    check(screen_size >= 3 and screen_size <= 9),
    description text,
    image text,
    price integer,
    check(price >= 0)
);