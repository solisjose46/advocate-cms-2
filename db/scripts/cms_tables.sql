CREATE TABLE Users (
    username TEXT,
    password TEXT,
    author_id INTEGER,
    PRIMARY KEY(username)
);

INSERT INTO Users (username, password, author_id)
    VALUES ("user1", "9fea8b6b39143e6e5338a16a6979b4bfaa76498d4cbe3467eb95b95bfaf54493", 1);