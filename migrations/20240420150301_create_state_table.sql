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