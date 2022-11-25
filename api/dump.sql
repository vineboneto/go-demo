DROP DATABASE go;

CREATE DATABASE go;

\CONNECT go;

CREATE TABLE IF NOT EXISTS tbl_usuario (
  id SERIAL PRIMARY KEY NOT NULL,
  email VARCHAR (255) NOT NULL UNIQUE,
  senha VARCHAR (255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL		
);

CREATE TABLE IF NOT EXISTS tbl_grupoacesso (
  id INT PRIMARY KEY NOT NULL,
  nome VARCHAR (255) NOT NULL UNIQUE		
);

CREATE TABLE IF NOT EXISTS tbl_grupoacesso_usuario (
  id_usuario INT NOT NULL,
  id_grupoacesso INT NOT NULL,

  PRIMARY KEY(id_grupoacesso, id_usuario),

  CONSTRAINT fk_usuario FOREIGN KEY(id_usuario) REFERENCES tbl_usuario(id),
  CONSTRAINT fk_grupoacesso FOREIGN KEY(id_grupoacesso) REFERENCES tbl_grupoacesso(id)
);

