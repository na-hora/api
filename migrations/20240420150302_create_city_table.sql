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
	CONSTRAINT "fk_city_state" FOREIGN KEY ("state_id") REFERENCES "public"."state" ("id") ON UPDATE CASCADE ON DELETE
	SET
		NULL
);