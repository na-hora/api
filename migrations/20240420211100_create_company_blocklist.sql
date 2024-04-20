-- Create "company_blocklist" table
CREATE TABLE "public"."company_blocklist" (
  "company_id" uuid NULL,
  "client_id" uuid NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "fk_company_blocklist_client" FOREIGN KEY ("client_id") REFERENCES "public"."client" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_blocklist_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
