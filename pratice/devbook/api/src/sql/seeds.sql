INSERT INTO users (name, nick, email, password)
VALUES
	('John Doe', 'jdoe', 'jdoe@gmail.com', '$2a$10$f/jjQsZjgQGWa56Sb6NL4u9CTzeWtuaIqMs//B1GVJAm.IxdMU7pO'),
	('Jane Doe', 'janed', 'janes@gmail.com', '$2a$10$f/jjQsZjgQGWa56Sb6NL4u9CTzeWtuaIqMs//B1GVJAm.IxdMU7pO'),
	('John Smith', 'jsmith', 'jsmith@gmail.com', '$2a$10$f/jjQsZjgQGWa56Sb6NL4u9CTzeWtuaIqMs//B1GVJAm.IxdMU7pO'),
	('Jane Smith', 'janes', 'janed@gmail.com', '$2a$10$f/jjQsZjgQGWa56Sb6NL4u9CTzeWtuaIqMs//B1GVJAm.IxdMU7pO');

INSERT INTO followers (user_id, follower_id) 
VALUES
	(1, 2),
	(1, 3),
	(2, 3),
	(3, 1);