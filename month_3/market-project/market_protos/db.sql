CREATE TYPE "st_tr_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "st_tr_source_type" AS ENUM (
  'sales',
  'bonus'
);

CREATE TYPE "br_pr_tr_type" AS ENUM (
  'plus',
  'minus'
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
  "updated_at" timestamp,
  "deleted_at" timestamp
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
  "updated_at" timestamp,
  "deleted_at" timestamp
);