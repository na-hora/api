-- Create "states" table
CREATE TABLE "public"."states" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
