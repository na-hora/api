-- Modify "company_pet_service_value" table
ALTER TABLE "public"."company_pet_service_value" DROP CONSTRAINT "fk_company_pet_service_value_company_pet_service", ADD
 CONSTRAINT "fk_company_pet_service_configurations" FOREIGN KEY ("company_pet_service_id") REFERENCES "public"."company_pet_service" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
