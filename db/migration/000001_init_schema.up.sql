CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" varchar,
  "date_of_birth" timestamptz,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "phone_number" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "product_types" (
  "id" bigserial PRIMARY KEY,
  "product_type_no" varchar NOT NULL,
  "title" varchar NOT NULL,
  "salient_features" text[],
  "descriptions" text[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "product_type_id" bigint NOT NULL,
  "price" bigint NOT NULL,
  "color" varchar,
  "size" bigint,
  "stock" bigint NOT NULL,
  "quantity_sold" bigint NOT NULL DEFAULT 0,
  "rating" bigint NOT NULL DEFAULT 0,
  "product_no" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "order_no" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "mt_product_order" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "order_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "cart_no" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "mt_product_cart" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "cart_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint,
  "discount" bigint
);

CREATE TABLE "warranties" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "title" varchar NOT NULL,
  "duration" bigint,
  "price" bigint,
  "times" bigint,
  "note" text[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "mt_product_type_warranty" (
  "id" bigserial PRIMARY KEY,
  "product_type_id" bigint NOT NULL,
  "warranty_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "url" varchar NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "mt_product_type_file" (
  "id" bigserial PRIMARY KEY,
  "purpose" varchar NOT NULL,
  "product_type_id" bigint NOT NULL,
  "file_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "content" text,
  "user_id" bigint NOT NULL,
  "product_type_id" bigint NOT NULL,
  "parent_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE TABLE "ratings" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "product_type_id" bigint NOT NULL,
  "stars" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" bigint,
  "updated_by" bigint
);

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "carts" ("user_id");

CREATE INDEX ON "mt_product_cart" ("product_id");

CREATE INDEX ON "mt_product_cart" ("cart_id");

CREATE INDEX ON "mt_product_cart" ("product_id", "cart_id");

COMMENT ON COLUMN "product_types"."discount" IS 'by %, gonna be overwrite by products.discount';

COMMENT ON COLUMN "products"."discount" IS 'by %, overwrite product_types.discount';

COMMENT ON COLUMN "orders"."discount" IS 'by %, store the discount at that moment';

COMMENT ON COLUMN "mt_product_order"."quantity" IS 'must be positive';

COMMENT ON COLUMN "mt_product_order"."total_price" IS 'store the total at that moment';

COMMENT ON COLUMN "mt_product_order"."discount" IS 'by %, store the discount at that moment';

COMMENT ON COLUMN "carts"."discount" IS 'by %, store the discount at that moment';

COMMENT ON COLUMN "mt_product_cart"."quantity" IS 'must be positive';

COMMENT ON COLUMN "mt_product_cart"."total_price" IS 'store the total at that moment';

COMMENT ON COLUMN "mt_product_cart"."discount" IS 'by %, store the discount at that moment';

COMMENT ON COLUMN "warranties"."type" IS 'can be free or paid';

COMMENT ON COLUMN "warranties"."duration" IS 'days';

COMMENT ON COLUMN "mt_product_type_file"."purpose" IS 'image for descriptions, image for thumnail';

COMMENT ON COLUMN "ratings"."stars" IS '1, 1.5, 2, 2.5 ... 5';

ALTER TABLE "products" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "mt_product_order" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "mt_product_order" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "mt_product_cart" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "mt_product_cart" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "mt_product_type_warranty" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");

ALTER TABLE "mt_product_type_warranty" ADD FOREIGN KEY ("warranty_id") REFERENCES "warranties" ("id");

ALTER TABLE "mt_product_type_file" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");

ALTER TABLE "mt_product_type_file" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");
