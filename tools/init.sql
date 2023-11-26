CREATE DATABASE project;

CREATE TABLE passwords (
    ID SERIAL PRIMARY KEY,
    password varchar(250) NOT NULL,
    strength INT NOT NULL
);

INSERT INTO passwords(password, strength) VALUES('intel1', 0);
INSERT INTO passwords(password, strength) VALUES('elyass15@ajilent-ci', 2);
INSERT INTO passwords(password, strength) VALUES('hodygid757#$!23w', 1);