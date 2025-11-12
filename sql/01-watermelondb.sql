DROP DATABASE IF EXISTS watermelon_db;
CREATE DATABASE watermelon_db;
USE watermelon_db;


DROP USER IF EXISTS watermelon;
CREATE USER 'watermelon'@'%' IDENTIFIED BY 'watermelon';
GRANT ALL PRIVILEGES ON watermelon_db.* TO 'watermelon'@'%';
FLUSH PRIVILEGES;


CREATE TABLE Usuario (
    Nome_de_usuario VARCHAR(200) NOT NULL,
    E_mail VARCHAR(100) UNIQUE NOT NULL,
    Senha VARCHAR(12) NOT NULL,
    Foto_de_Perfil MEDIUMBLOB,
    ID_usuario INT PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE Relatorio (
    Tipo_de_relatorio ENUM('simples', 'avancado') NOT NULL,
    Data DATE NOT NULL,
    Descricao LONGTEXT NOT NULL,
    Temperatura DECIMAL(4, 2),
    Turbidez DECIMAL(3, 2),
    Cor_Aparente VARCHAR(36) NOT NULL,
    Acidez DECIMAL(2, 1),
    Oxigenio_Dissolvido DECIMAL(3, 3),
    Demanda_Bioquimica_de_Oxigenio DECIMAL(3, 3),
    Nitrogenio_Total DECIMAL(3, 3),
    Fosforo_Total DECIMAL(3, 3),
    Metais_Pesados DECIMAL(5, 3),
    Cloro_Residual DECIMAL(3, 3),
    Composto_Organico_Volatil DECIMAL(4, 3),
    Coliformes DECIMAL(3, 3),
    Avaliacao_Biologica DECIMAL(3, 3),
    Solidos_Totais_Dissolvidos DECIMAL(3, 3),
    Solidos_em_Suspensao DECIMAL(3, 3),
    Odor VARCHAR(100),
    Sabor VARCHAR(100),
    ID_relatorio INT PRIMARY KEY AUTO_INCREMENT,
    ID_Autor INT NOT NULL,
    ID_Corpo_d_agua INT NOT NULL
);

CREATE TABLE Corpo_d_agua (
    Nome VARCHAR(100) NOT NULL,
    Tipo ENUM('abra', 'angra', 'arroio', 'bacia', 'baia', 'barragem', 'canal', 'corredeira', 'corrego', 'correnteza', 'enseada', 'estreito', 'estuario', 'fiorde', 'fitotelmo', 'fosso', 'geleira', 'golfo', 'lago_subglacial', 'lago', 'lagoa', 'laguna', 'mar', 'marisma', 'nascente', 'oceano', 'regato', 'reservatorio', 'rio', 'riacho', 'riachuelo', 'ribeira', 'ribeirao', 'sanga', 'zona_umida') NOT NULL,
    ID_Corpo_d_agua INT PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE Localizacao (
    Cidade VARCHAR(50) NOT NULL,
    Estado ENUM('ac', 'al', 'am', 'ap', 'ba', 'ce', 'df', 'es', 'go', 'ma', 'mg', 'ms', 'mt', 'pa', 'pb', 'pe', 'pi', 'pr', 'rj', 'rn', 'ro', 'rr', 'rs', 'sc', 'se', 'sp', 'to') NOT NULL,
    ID_localizacao INT PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE Corpo_Localizacao (
    ID_corpo_localizacao INT PRIMARY KEY AUTO_INCREMENT,
    ID_Corpo_d_agua INT NOT NULL,
    ID_Localizacao INT NOT NULL
);
 
ALTER TABLE Relatorio ADD CONSTRAINT FK_Relatorio_2
    FOREIGN KEY (ID_Autor)
    REFERENCES Usuario (ID_usuario)
    ON DELETE CASCADE;
 
ALTER TABLE Relatorio ADD CONSTRAINT FK_Relatorio_3
    FOREIGN KEY (ID_Corpo_d_agua)
    REFERENCES Corpo_d_agua (ID_Corpo_d_agua)
    ON DELETE CASCADE;
 
ALTER TABLE Corpo_Localizacao ADD CONSTRAINT FK_Corpo_Localizacao_2
    FOREIGN KEY (ID_Corpo_d_agua)
    REFERENCES Corpo_d_agua (ID_Corpo_d_agua);
 
ALTER TABLE Corpo_Localizacao ADD CONSTRAINT FK_Corpo_Localizacao_3
    FOREIGN KEY (ID_Localizacao)
    REFERENCES Localizacao (ID_localizacao);
