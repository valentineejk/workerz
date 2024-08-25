-- Add up migration script here
CREATE TABLE "workerz" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" varchar NOT NULL,
  "country" varchar NOT NULL,
  "salary" float NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);