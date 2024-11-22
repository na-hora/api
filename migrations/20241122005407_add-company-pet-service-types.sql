-- Create "company_pet_service_types" table
CREATE TABLE "public"."company_pet_service_types" (
  "company_pet_service_id" bigint NOT NULL,
  "company_pet_type_id" integer NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("company_pet_service_id", "company_pet_type_id"),
  CONSTRAINT "fk_company_pet_service_service_types" FOREIGN KEY ("company_pet_service_id") REFERENCES "public"."company_pet_service" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_pet_type_service_types" FOREIGN KEY ("company_pet_type_id") REFERENCES "public"."company_pet_type" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_pet_service_types_deleted_at" to table: "company_pet_service_types"
CREATE INDEX "idx_company_pet_service_types_deleted_at" ON "public"."company_pet_service_types" ("deleted_at");
