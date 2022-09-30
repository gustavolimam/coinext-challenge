-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "name" varchar,
    "age" int2,
    "gender" bpchar(1),
    "latitude" varchar,
    "longitude" varchar,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "dead_at" timestamp,
    PRIMARY KEY ("id")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS inventory_user_id_seq;

-- Table Definition
CREATE TABLE "inventory" (
    "user_id" int4 NOT NULL DEFAULT nextval('inventory_user_id_seq'::regclass),
    "water" int2,
    "food" int2,
    "drug" int2,
    "ammo" int2,
    CONSTRAINT "fk_customer" FOREIGN KEY ("user_id") REFERENCES "users"("id"),
    PRIMARY KEY ("user_id")
);