CREATE TABLE Account (
    IdUser SERIAL PRIMARY KEY,
    Description VARCHAR,
    Mail VARCHAR NOT NULL,

    UNIQUE(mail)
);

CREATE TABLE InfoSecurAccount (
    IdUser INTEGER  NOT NULL,
    Username VARCHAR(25) NOT NULL,
    Password VARCHAR NOT NULL,
    VerificationWord VARCHAR(15) NOT NULL,

    FOREIGN KEY(IdUser) REFERENCES Account(IdUser)
);

CREATE TABLE Meme (
    IdMeme SERIAL PRIMARY KEY,
    IdUser INTEGER  NOT NULL,
    MemeName VARCHAR(25) NOT NULL,
    UrlMeme VARCHAR NOT NULL,
    Description VARCHAR,
    Topic VARCHAR,
    Favorite INTEGER NOT NULL,
    
    FOREIGN KEY(IdUser) REFERENCES Account(IdUser)
);