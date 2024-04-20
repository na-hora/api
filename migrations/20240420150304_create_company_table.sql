-- Create "company" table
CREATE TABLE "public"."company" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"name" text NOT NULL,
	"fantasy_name" text NOT NULL,
	"cnpj" text NOT NULL,
	"email" text NOT NULL,
	"phone" text NOT NULL,
	"avatar_url" text NULL,
	"category_id" integer NOT NULL,
	"created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"deleted_at" timestamptz NULL,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_company_category" FOREIGN KEY ("category_id") REFERENCES "public"."company_category" ("id") ON UPDATE CASCADE ON DELETE
	SET
		NULL
);

-- Create index "idx_company_deleted_at" to table: "company"
CREATE INDEX "idx_company_deleted_at" ON "public"."company" ("deleted_at");