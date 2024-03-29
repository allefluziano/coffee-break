CREATE DATABASE IF NOT EXISTS coffeebreak;
USE coffeebreak;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp() on UPDATE current_timestamp()
) ENGINE=INNODB;
