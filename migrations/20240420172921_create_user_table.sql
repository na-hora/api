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
