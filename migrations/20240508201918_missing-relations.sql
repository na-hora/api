-- Modify "appointment" table
ALTER TABLE "public"."appointment" DROP CONSTRAINT "fk_appointment_client", ADD
 CONSTRAINT "fk_client_appointments" FOREIGN KEY ("client_id") REFERENCES "public"."client" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
