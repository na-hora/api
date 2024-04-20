-- Create "pet_size" table
CREATE TABLE "public"."pet_size" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "uni_pet_size_name" to table: "pet_size"
CREATE UNIQUE INDEX "uni_pet_size_name" ON "public"."pet_size" ("name");
