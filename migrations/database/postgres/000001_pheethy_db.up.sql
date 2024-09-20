-- Create transaction --
BEGIN;

-- set time zone --
SET TIME ZONE 'Asia/Bangkok';

-- install Extension UUID --
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create enum --
CREATE TYPE order_status AS ENUM (
    'WAITING',
    'SHIPPING',
    'COMPLETED',
    'CANCELED'
);

CREATE TABLE "products" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" VARCHAR NOT NULL UNIQUE,
  "description" VARCHAR NOT NULL DEFAULT '',
  "price" FLOAT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "products_categories" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "product_id" uuid NOT NULL,
  "category_id" INT NOT NULL
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "users" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "role_id" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "oauth" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "access_token" VARCHAR NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR NOT NULL
);

CREATE TABLE "images" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "filename" VARCHAR NOT NULL,
  "url" VARCHAR NOT NULL,
  "product_id" uuid NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "orders" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "contact" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "status" order_status NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "transfer_slip" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "order_id" uuid NOT NULL UNIQUE,
  "filename" VARCHAR NOT NULL,
  "url" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "products_orders" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "order_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "qty" INT NOT NULL DEFAULT 1
);

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
ALTER TABLE "products_orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "transfer_slip" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

COMMIT;