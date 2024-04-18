-- Create "cities" table
CREATE TABLE "public"."cities" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "state_id" integer NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_cities_state" FOREIGN KEY ("state_id") REFERENCES "public"."states" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
