CREATE TYPE "staff_type" AS ENUM (
  'cashier',
  'shop_assistant'
);

CREATE TYPE "tarif_type" AS ENUM (
  'fixed',
  'percent'
);

CREATE TYPE "st_tr_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "st_tr_source_type" AS ENUM (
  'sales',
  'bonus'
);

CREATE TYPE "sale_payment_type" AS ENUM (
  'card',
  'cash'
);

CREATE TYPE "sale_status" AS ENUM (
  'success',
  'cancel'
);

CREATE TYPE "br_pr_tr_type" AS ENUM (
  'plus',
  'minus'
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "parent_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" numeric NOT NULL,
  "barcode" varchar NOT NULL,
  "category_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "address" varchar,CREATE TYPE "staff_type" AS ENUM (
  'cashier',
  'shop_assistant'
);

CREATE TYPE "tarif_type" AS ENUM (
  'fixed',
  'percent'
);

CREATE TYPE "st_tr_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "st_tr_source_type" AS ENUM (
  'sales',
  'bonus'
);

CREATE TYPE "sale_payment_type" AS ENUM (
  'card',
  'cash'
);

CREATE TYPE "sale_status" AS ENUM (
  'success',
  'cancel'
);

CREATE TYPE "br_pr_tr_type" AS ENUM (
  'plus',
  'minus'
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "parent_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" numeric NOT NULL,
  "barcode" varchar NOT NULL,
  "category_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "address" varchar,
  "founded_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "branch_products" (
  "product_id" uuid,
  "branch_id" uuid,
  "count" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff_tarif" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "type" tarif_type,
  "amount_for_cash" numeric,
  "amount_for_card" numeric,
  "founded_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "tarif_id" uuid,
  "name" varchar,
  "type" staff_type,
  "balance" numeric,
  "birth_date" date,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff_transactions" (
  "id" uuid PRIMARY KEY,
  "sale_id" uuid,
  "staff_id" uuid,
  "type" st_tr_type,
  "source_type" st_tr_source_type,
  "amount" numeric,
  "about_text" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "sales" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "shop_assistant_id" uuid,
  "cashier_id" uuid,
  "price" numeric,
  "payment_type" sale_payment_type,
  "status" sale_status,
  "client_name" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "sale_products" (
  "sale_id" uuid,
  "product_id" uuid,
  "quantity" integer,
  "price" numeric
);

CREATE TABLE "branch_product_transactions" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "staff_id" uuid,
  "product_id" uuid,
  "price" numeric,
  "type" br_pr_tr_type,
  "quantity" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("tarif_id") REFERENCES "staff_tarif" ("id");

ALTER TABLE "branch_products" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "staff_transactions" ADD FOREIGN KEY ("sale_id") REFERENCES "sales" ("id");

ALTER TABLE "sale_products" ADD FOREIGN KEY ("sale_id") REFERENCES "sales" ("id");
  "founded_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "branch_products" (
  "product_id" uuid,
  "branch_id" uuid,
  "count" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff_tarif" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "type" tarif_type,
  "amount_for_cash" numeric,
  "amount_for_card" numeric,
  "founded_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "tarif_id" uuid,
  "name" varchar,
  "type" staff_type,
  "balance" numeric,
  "birth_date" date,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "staff_transactions" (
  "id" uuid PRIMARY KEY,
  "sale_id" uuid,
  "staff_id" uuid,
  "type" st_tr_type,
  "source_type" st_tr_source_type,
  "amount" numeric,
  "about_text" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "sales" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "shop_assistant_id" uuid,
  "cashier_id" uuid,
  "price" numeric,
  "payment_type" sale_payment_type,
  "status" sale_status,
  "client_name" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "sale_products" (
  "sale_id" uuid,
  "product_id" uuid,
  "quantity" integer,
  "price" numeric
);

CREATE TABLE "branch_product_transactions" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "staff_id" uuid,
  "product_id" uuid,
  "price" numeric,
  "type" br_pr_tr_type,
  "quantity" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("tarif_id") REFERENCES "staff_tarif" ("id");

ALTER TABLE "branch_products" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "staff_transactions" ADD FOREIGN KEY ("sale_id") REFERENCES "sales" ("id");

ALTER TABLE "sale_products" ADD FOREIGN KEY ("sale_id") REFERENCES "sales" ("id");