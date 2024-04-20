-- Create "company_category" table
CREATE TABLE "public"."company_category" (
	"id" bigserial NOT NULL,
	"name" text NOT NULL,
	"created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"deleted_at" timestamptz NULL,
	PRIMARY KEY ("id")
);

-- Create index "idx_company_category_deleted_at" to table: "company_category"
CREATE INDEX "idx_company_category_deleted_at" ON "public"."company_category" ("deleted_at");