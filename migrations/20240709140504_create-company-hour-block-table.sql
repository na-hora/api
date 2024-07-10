-- Create "company_hour_block" table
CREATE TABLE "public"."company_hour_block" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "day" timestamptz NOT NULL,
  "start_hour" bigint NOT NULL,
  "end_hour" bigint NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_hour_block_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_hour_block_deleted_at" to table: "company_hour_block"
CREATE INDEX "idx_company_hour_block_deleted_at" ON "public"."company_hour_block" ("deleted_at");
