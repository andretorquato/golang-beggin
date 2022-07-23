CREATE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id int auto_increment primary key,
	username varchar(50) not null,
	nickname varchar(50) not null,
	email varchar(50) not null,
	password varchar(50) not null,
	created_at timestamp default current_timestamp(),
) ENGINE=INNODB;

