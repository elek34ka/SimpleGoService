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

DROP TABLE IF EXISTS user_in_group;
CREATE TABLE user_in_group (
        user_id INT NOT NULL REFERENCES users(id),
        group_id INT NOT NULL REFERENCES `groups`(id),
        UNIQUE (user_id, group_id)
);
