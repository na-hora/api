-- Create "token" table
CREATE TABLE "public"."token" (
  "key" uuid NOT NULL DEFAULT gen_random_uuid(),
  "note" text NULL,
  "company_id" uuid NULL,
  "used" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("key"),
  CONSTRAINT "fk_token_company" FOREIGN KEY ("company_id") REFERENCES "public"."company" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
