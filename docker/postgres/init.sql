CREATE TABLE users (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
email VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL
);

INSERT INTO users (name, email, password) VALUES
('John Doe', 'john@example.com', 'securepassword'),
('toby nguyen', 'toby@example.com', 'anotherpassword'),
('toby nguyen1', 'toby@example.com1', 'yetanotherpassword'),
('toby nguyen2', 'toby@example.com2', 'yetanothercoolpassword'),
('toby nguyen3', 'toby@example.com3', 'yetanothercoolamazingpassword'),
('toby nguyen4', 'toby@example.com4', 'yeah'),
('toby nguyen5', 'toby@example.com5', 'yeahh');