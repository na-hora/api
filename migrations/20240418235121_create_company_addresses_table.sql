-- Create "company_addresses" table
CREATE TABLE "public"."company_addresses" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "company_id" uuid NOT NULL,
  "zip_code" text NULL,
  "city_id" integer NULL,
  "neighborhood" text NULL,
  "street" text NULL,
  "number" text NULL,
  "complement" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_addresses_city" FOREIGN KEY ("city_id") REFERENCES "public"."cities" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_addresses_company" FOREIGN KEY ("company_id") REFERENCES "public"."companies" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
