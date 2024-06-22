-- Modify "token" table
ALTER TABLE "public"."token" ADD COLUMN "user_id" uuid NULL, ADD
 CONSTRAINT "fk_token_user" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
