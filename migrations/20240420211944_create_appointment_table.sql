-- Create "appointment" table
CREATE TABLE "public"."appointment" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "company_id" uuid NOT NULL,
  "client_id" uuid NOT NULL,
  "pet_breed_id" integer NOT NULL,
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
  CONSTRAINT "fk_appointment_client" FOREIGN KEY ("client_id") REFERENCES "public"."client" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_appointment_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_appointment_pet_breed" FOREIGN KEY ("pet_breed_id") REFERENCES "public"."pet_breed" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
