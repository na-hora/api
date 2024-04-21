-- Modify "appointment" table
ALTER TABLE "public"."appointment" DROP CONSTRAINT "fk_appointment_company", ADD
 CONSTRAINT "fk_company_appointments" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
-- Modify "company_blocklist" table
ALTER TABLE "public"."company_blocklist" DROP CONSTRAINT "fk_company_blocklist_company", ADD
 CONSTRAINT "fk_company_blocks" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
-- Modify "company_pet_breed" table
ALTER TABLE "public"."company_pet_breed" DROP CONSTRAINT "fk_company_pet_breed_company", ADD
 CONSTRAINT "fk_company_company_pet_breeds" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
-- Modify "company_pet_size" table
ALTER TABLE "public"."company_pet_size" DROP CONSTRAINT "fk_company_pet_size_company", ADD
 CONSTRAINT "fk_company_company_pet_sizes" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
