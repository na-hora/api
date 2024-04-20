-- Create "company_pet_breed" table
CREATE TABLE "public"."company_pet_breed" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "pet_breed_id" integer NOT NULL,
  "extra_value" numeric NULL,
  "extra_time" numeric NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_pet_breed_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_pet_breed_pet_breed" FOREIGN KEY ("pet_breed_id") REFERENCES "public"."pet_breed" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
