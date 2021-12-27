DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    debt INT NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS userInGroup;
CREATE TABLE userInGroup (
        id INT AUTO_INCREMENT NOT NULL,
        user_id INT NOT NULL REFERENCES users(id),
        group_id INT NOT NULL REFERENCES `groups`(id),
        PRIMARY KEY (id)
);
