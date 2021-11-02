CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS restaurants(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), name VARCHAR (100) not null, address VARCHAR (100), city VARCHAR (50));
CREATE TABLE IF NOT EXISTS menu_items(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), name VARCHAR (50) not null, category VARCHAR(50), cuisine VARCHAR (50), price NUMERIC, restaurant_id uuid);
ALTER TABLE menu_items ADD CONSTRAINT fk_menu_items_restaurants FOREIGN KEY (restaurant_id) REFERENCES restaurants(id);