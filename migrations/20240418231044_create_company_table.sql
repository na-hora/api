-- Create "companies" table
CREATE TABLE "public"."companies" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "fantasy_name" text NOT NULL,
  "cnpj" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NOT NULL,
  "avatar_url" text NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_companies_deleted_at" to table: "companies"
CREATE INDEX "idx_companies_deleted_at" ON "public"."companies" ("deleted_at");
