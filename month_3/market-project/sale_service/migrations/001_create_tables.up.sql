CREATE TYPE "status_type" AS ENUM (
  'success',
  'cancel'
);

CREATE TYPE "payment_type" AS ENUM (
  'card',
  'cash'
);

CREATE TYPE "transaction_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "source_type" AS ENUM (
  'sales',
  'bonus'
);

CREATE TABLE "sale" (
  "id" uuid PRIMARY KEY,
  "client_name" varchar(20) NOT NULL,
  "branch_id" uuid NOT NULL,
  "cashier_id" uuid NOT NULL,
  "shop_assistant_id" uuid,
  "price" NUMERIC(12, 2),
  "status" status_type DEFAULT 'success',
  "payment_type" payment_type NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "sale_products" (
  "id" uuid PRIMARY KEY,
  "sale_id" uuid NOT NULL REFERENCES "sale"("id"),
  "product_id" uuid NOT NULL,
  "quantity" integer,
  "price" NUMERIC(12, 2),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "staff_transactions" (
  "id" uuid PRIMARY KEY,
  "type" transaction_type NOT NULL,
  "amount" NUMERIC(12, 2),
  "source_type" source_type NOT NULL,
  "sale_id" uuid NOT NULL REFERENCES "sale"("id"),
  "staff_id" uuid NOT NULL,
  "text" varchar(100),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);
