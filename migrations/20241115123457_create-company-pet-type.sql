-- Create "company_pet_type" table
CREATE TABLE "public"."company_pet_type" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_pet_type_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Modify "company_pet_hair" table
ALTER TABLE "public"."company_pet_hair" ADD COLUMN "company_pet_type_id" integer NOT NULL, ADD
 CONSTRAINT "fk_company_pet_hair_company_pet_type" FOREIGN KEY ("company_pet_type_id") REFERENCES "public"."company_pet_type" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
-- Modify "company_pet_size" table
ALTER TABLE "public"."company_pet_size" ADD COLUMN "company_pet_type_id" integer NOT NULL, ADD
 CONSTRAINT "fk_company_pet_size_company_pet_type" FOREIGN KEY ("company_pet_type_id") REFERENCES "public"."company_pet_type" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
