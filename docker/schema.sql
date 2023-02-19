USE pessoas;

CREATE TABLE pessoa (
    id INT NOT NULL AUTO_INCREMENT,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    idade INT(2) NOT NULL,
    email VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
)