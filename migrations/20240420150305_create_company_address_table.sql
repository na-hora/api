
-- Create "company_address" table
CREATE TABLE "public"."company_address" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"company_id" uuid NOT NULL,
	"zip_code" text NULL,
	"city_id" bigint NULL,
	"neighborhood" text NULL,
	"street" text NULL,
	"number" text NULL,
	"complement" text NULL,
	"created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_company_address_city" FOREIGN KEY ("city_id") REFERENCES "public"."city" ("id") ON UPDATE CASCADE ON DELETE
	SET
		NULL,
		CONSTRAINT "fk_company_address_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE
	SET
		NULL
);