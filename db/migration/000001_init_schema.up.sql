CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" varchar,
  "date_of_birth" timestamptz,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "phone_number" varchar
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "size" bigint,
  "color" varchar,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "order_no" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "mt_product_order" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "order_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "mt_product_order" ("product_id");

CREATE INDEX ON "mt_product_order" ("order_id");

CREATE INDEX ON "mt_product_order" ("product_id", "order_id");

COMMENT ON COLUMN "mt_product_order"."quantity" IS 'must be positive';

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "mt_product_order" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "mt_product_order" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
