-- Create "company_category" table
CREATE TABLE "public"."company_category" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_company_category_deleted_at" to table: "company_category"
CREATE INDEX "idx_company_category_deleted_at" ON "public"."company_category" ("deleted_at");
-- Create "company" table
CREATE TABLE "public"."company" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "fantasy_name" text NOT NULL,
  "cnpj" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NOT NULL,
  "avatar_url" text NULL,
  "category_id" integer NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_category" FOREIGN KEY ("category_id") REFERENCES "public"."company_category" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_deleted_at" to table: "company"
CREATE INDEX "idx_company_deleted_at" ON "public"."company" ("deleted_at");
-- Create "company_pet_hair" table
CREATE TABLE "public"."company_pet_hair" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_company_pet_hairs" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create "company_pet_service" table
CREATE TABLE "public"."company_pet_service" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "name" text NOT NULL,
  "concurrency" bigint NOT NULL DEFAULT 1,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_company_pet_services" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_pet_service_deleted_at" to table: "company_pet_service"
CREATE INDEX "idx_company_pet_service_deleted_at" ON "public"."company_pet_service" ("deleted_at");
-- Create "company_pet_size" table
CREATE TABLE "public"."company_pet_size" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_company_pet_sizes" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create "company_pet_service_value" table
CREATE TABLE "public"."company_pet_service_value" (
  "id" bigserial NOT NULL,
  "company_pet_service_id" bigint NOT NULL,
  "company_pet_size_id" integer NOT NULL,
  "company_pet_hair_id" integer NOT NULL,
  "price" numeric NOT NULL,
  "execution_time" bigint NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_pet_service_value_company_pet_hair" FOREIGN KEY ("company_pet_hair_id") REFERENCES "public"."company_pet_hair" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_pet_service_value_company_pet_service" FOREIGN KEY ("company_pet_service_id") REFERENCES "public"."company_pet_service" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_pet_service_value_company_pet_size" FOREIGN KEY ("company_pet_size_id") REFERENCES "public"."company_pet_size" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_pet_service_value_deleted_at" to table: "company_pet_service_value"
CREATE INDEX "idx_company_pet_service_value_deleted_at" ON "public"."company_pet_service_value" ("deleted_at");
-- Create "client" table
CREATE TABLE "public"."client" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "phone" text NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create "appointment" table
CREATE TABLE "public"."appointment" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "company_id" uuid NOT NULL,
  "client_id" uuid NOT NULL,
  "company_pet_service_value_id" bigint NOT NULL,
  "pet_name" text NULL,
  "start_time" timestamptz NOT NULL,
  "total_time" bigint NOT NULL,
  "total_price" numeric NOT NULL,
  "payment_mode" text NULL,
  "canceled" boolean NULL,
  "cancelation_reason" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_appointment_company_pet_service_value" FOREIGN KEY ("company_pet_service_value_id") REFERENCES "public"."company_pet_service_value" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_client_appointments" FOREIGN KEY ("client_id") REFERENCES "public"."client" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_appointments" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create "state" table
CREATE TABLE "public"."state" (
  "id" bigserial NOT NULL,
  "uf" text NOT NULL,
  "name" text NOT NULL,
  "ibge" bigint NOT NULL,
  "ddd" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "uni_state_name" to table: "state"
CREATE UNIQUE INDEX "uni_state_name" ON "public"."state" ("name");
-- Create index "uni_state_uf" to table: "state"
CREATE UNIQUE INDEX "uni_state_uf" ON "public"."state" ("uf");
-- Create "city" table
CREATE TABLE "public"."city" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "state_id" bigint NOT NULL,
  "ibge" text NOT NULL,
  "lat_lon" text NULL,
  "cod_tom" bigint NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_city_state" FOREIGN KEY ("state_id") REFERENCES "public"."state" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create "company_address" table
CREATE TABLE "public"."company_address" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "company_id" uuid NOT NULL,
  "zip_code" text NULL,
  "city_id" bigint NULL,
  "neighborhood" text NULL,
  "street" text NULL,
  "number" bigint NULL,
  "complement" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_address_city" FOREIGN KEY ("city_id") REFERENCES "public"."city" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_company_company_addresses" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create "company_hour" table
CREATE TABLE "public"."company_hour" (
  "id" bigserial NOT NULL,
  "company_id" uuid NOT NULL,
  "weekday" bigint NOT NULL,
  "start_hour" bigint NOT NULL,
  "end_hour" bigint NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_company_hours" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_company_hour_deleted_at" to table: "company_hour"
CREATE INDEX "idx_company_hour_deleted_at" ON "public"."company_hour" ("deleted_at");
-- Create "user" table
CREATE TABLE "public"."user" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "company_id" uuid NOT NULL,
  "username" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_company_users" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "uni_user_username" to table: "user"
CREATE UNIQUE INDEX "uni_user_username" ON "public"."user" ("username");
