-- Modify "token" table
ALTER TABLE "public"."token" ADD COLUMN "expires_at" timestamptz NULL;
