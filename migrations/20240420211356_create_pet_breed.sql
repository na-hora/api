-- Create "pet_breed" table
CREATE TABLE "public"."pet_breed" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "type" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "uni_pet_breed_name" to table: "pet_breed"
CREATE UNIQUE INDEX "uni_pet_breed_name" ON "public"."pet_breed" ("name");
