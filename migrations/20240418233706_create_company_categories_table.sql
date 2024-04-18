-- Create "company_categories" table
CREATE TABLE "public"."company_categories" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_company_categories_deleted_at" to table: "company_categories"
CREATE INDEX "idx_company_categories_deleted_at" ON "public"."company_categories" ("deleted_at");
-- Modify "companies" table
ALTER TABLE "public"."companies" ALTER COLUMN "category_id" TYPE integer, ADD
 CONSTRAINT "fk_companies_category" FOREIGN KEY ("category_id") REFERENCES "public"."company_categories" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
