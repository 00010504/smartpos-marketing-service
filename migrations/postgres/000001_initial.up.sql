CREATE TYPE product_type AS ENUM ('8b0bf29c-58e8-4310-8bb1-a1b9771f9c47','2b98f424-91c9-46cc-abd7-c888208807da', 'a19a514e-41c9-4666-a01a-e3f9c0255609');
CREATE TYPE custom_field_type AS ENUM ('8b0bf29c-58e8-4310-8bb1-a1b9771f9c47','2b98f424-91c9-46cc-abd7-c888208807da', 'a19a514e-41c9-4666-a01a-e3f9c0255609');


CREATE TYPE user_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- system user
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin user
);

CREATE TYPE app_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- client
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin
);

CREATE TABLE "sex" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB
);

INSERT INTO "sex" (
    "id",
    "name"
) VALUES (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    'male'
), (
    '9fb3ada6-a73b-4b81-9295-5c1605e54552',
    'Female'
);

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY,
    "user_type_id" user_type NOT NULL,
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "phone_number" VARCHAR(30) NOT NULL,
    "image" TEXT,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX "user_deleted_at_idx" ON "user"("deleted_at");

INSERT INTO "user" (
    "id",
    "first_name",
    "last_name",
    "phone_number",
    "user_type_id"
) VALUES (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    'admin',
    'admin',
    '99894172774',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'
);


CREATE TABLE IF NOT EXISTS "company" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL
);
CREATE INDEX company_deleted_at_idx ON "company"("deleted_at");

CREATE TABLE IF NOT EXISTS "company_user" (
    "user_id" UUID NOT NULL REFERENCES "user" ("id"),
    "company_id" UUID NOT NULL,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY("user_id", "company_id", "deleted_at")
);

CREATE INDEX "company_user_deleted_at_idx" ON "company_user"("deleted_at");


CREATE TABLE IF NOT EXISTS "shop" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "company_id" UUID NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX shop_deleted_at_idx ON "shop"("deleted_at");


CREATE TABLE IF NOT EXISTS "group" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    UNIQUE("company_id", "name", "deleted_at")
);


CREATE TABLE IF NOT EXISTS "client" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "sex_id" UUID REFERENCES "sex"("id") ON DELETE SET NULL,
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "email" VARCHAR(100),
    "phone_number" VARCHAR,
    "info" TEXT,
    "birthday" DATE,
    "created_by" UUID,
    "card_number" VARCHAR,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID
);

CREATE TABLE IF NOT EXISTS "client_group" (
    "client_id" UUID NOT NULL,
    "group_id" UUID NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    PRIMARY KEY ("client_id", "group_id", "deleted_at")
);