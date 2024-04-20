-- Create "company_hour" table
CREATE TABLE "public"."company_hour" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "weekday" bigint NOT NULL,
  "start_hour" numeric NOT NULL,
  "end_hour" numeric NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_hour_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_hour_deleted_at" to table: "company_hour"
CREATE INDEX "idx_company_hour_deleted_at" ON "public"."company_hour" ("deleted_at");
