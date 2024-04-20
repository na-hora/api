-- Create "company_pet_size" table
CREATE TABLE "public"."company_pet_size" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "pet_size_id" integer NOT NULL,
  "extra_value" numeric NOT NULL,
  "extra_time" numeric NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_pet_size_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_pet_size_pet_size" FOREIGN KEY ("pet_size_id") REFERENCES "public"."pet_size" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
