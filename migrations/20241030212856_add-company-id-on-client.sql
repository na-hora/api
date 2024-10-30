-- Modify "client" table
ALTER TABLE "public"."client" ADD COLUMN "company_id" uuid NOT NULL, ADD
 CONSTRAINT "fk_client_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
