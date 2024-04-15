CREATE DATABASE IF NOT EXISTS proyecto2;

USE proyecto2;

CREATE TABLE IF NOT EXISTS votos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name_v VARCHAR(50),
    album_v VARCHAR(50),
    year_v INT,
    rank_v INT
);