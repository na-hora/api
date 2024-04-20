-- Modify "cities" table
ALTER TABLE "public"."cities" ALTER COLUMN "state_id" TYPE bigint, ADD COLUMN "ibge" text NOT NULL, ADD COLUMN "lat_lon" text NULL, ADD COLUMN "cod_tom" bigint NULL;
-- Modify "company_addresses" table
ALTER TABLE "public"."company_addresses" ALTER COLUMN "city_id" TYPE bigint;
-- Modify "states" table
ALTER TABLE "public"."states" ADD COLUMN "uf" text NOT NULL, ADD COLUMN "ibge" bigint NOT NULL, ADD COLUMN "ddd" text NULL;
-- Create index "uni_states_name" to table: "states"
CREATE UNIQUE INDEX "uni_states_name" ON "public"."states" ("name");
-- Create index "uni_states_uf" to table: "states"
CREATE UNIQUE INDEX "uni_states_uf" ON "public"."states" ("uf");
