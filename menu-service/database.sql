CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS restaurants(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), name VARCHAR (100) not null, address VARCHAR (100), city VARCHAR (50));