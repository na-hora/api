-- Create "service" table
CREATE TABLE "public"."service" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "name" text NOT NULL,
  "price" numeric NOT NULL,
  "execution_time" bigint NOT NULL,
  "concurrency" bigint NOT NULL DEFAULT 1,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_services" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_service_deleted_at" to table: "service"
CREATE INDEX "idx_service_deleted_at" ON "public"."service" ("deleted_at");
