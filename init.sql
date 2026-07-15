CREATE TABLE "contacts"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "nama" VARCHAR(100) NOT NULL, 
    "email" VARCHAR(100) NOT NULL UNIQUE,
    "phone" VARCHAR(100) NOT NULL UNIQUE,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

INSERT INTO "contacts" ("nama", "email", "phone")
VALUES
('pra', 'pra@gmail.com', '08123456789'),
('Dimas', 'dimas@gmail.com', '08111111111');

SELECT "nama", "email", "phone" FROM "contacts";
