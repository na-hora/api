-- Modify "company_address" table
ALTER TABLE "public"."company_address" DROP CONSTRAINT "fk_company_address_company", ADD
 CONSTRAINT "fk_company_company_addresses" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
-- Modify "company_hour" table
ALTER TABLE "public"."company_hour" DROP CONSTRAINT "fk_company_hour_company", ADD
 CONSTRAINT "fk_company_company_hours" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
