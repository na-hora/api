-- Modify "company_pet_hair" table
ALTER TABLE "public"."company_pet_hair" ADD COLUMN "deleted_at" timestamptz NULL;
-- Create index "idx_company_pet_hair_deleted_at" to table: "company_pet_hair"
CREATE INDEX "idx_company_pet_hair_deleted_at" ON "public"."company_pet_hair" ("deleted_at");
-- Modify "company_pet_size" table
ALTER TABLE "public"."company_pet_size" ADD COLUMN "deleted_at" timestamptz NULL;
-- Create index "idx_company_pet_size_deleted_at" to table: "company_pet_size"
CREATE INDEX "idx_company_pet_size_deleted_at" ON "public"."company_pet_size" ("deleted_at");
-- Modify "company_pet_type" table
ALTER TABLE "public"."company_pet_type" ADD COLUMN "deleted_at" timestamptz NULL;
-- Create index "idx_company_pet_type_deleted_at" to table: "company_pet_type"
CREATE INDEX "idx_company_pet_type_deleted_at" ON "public"."company_pet_type" ("deleted_at");
