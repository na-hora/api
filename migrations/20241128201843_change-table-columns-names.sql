-- Rename a column from "start_hour" to "start_minute"
ALTER TABLE "public"."company_hour" RENAME COLUMN "start_hour" TO "start_minute";
-- Rename a column from "end_hour" to "end_minute"
ALTER TABLE "public"."company_hour" RENAME COLUMN "end_hour" TO "end_minute";
