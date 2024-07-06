-- Rename a column from "concurrency" to "paralellism"
ALTER TABLE "public"."company_pet_service" RENAME COLUMN "concurrency" TO "paralellism";
-- Modify "state" table
ALTER TABLE "public"."state" ADD CONSTRAINT "uni_state_name" UNIQUE USING INDEX "uni_state_name", ADD CONSTRAINT "uni_state_uf" UNIQUE USING INDEX "uni_state_uf";
-- Modify "user" table
ALTER TABLE "public"."user" ADD CONSTRAINT "uni_user_username" UNIQUE USING INDEX "uni_user_username";
