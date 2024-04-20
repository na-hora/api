-- Modify "cities" table
ALTER TABLE
	"public"."cities"
ALTER COLUMN
	"created_at"
SET
	DEFAULT now();

-- Modify "companies" table
ALTER TABLE
	"public"."companies"
ALTER COLUMN
	"created_at"
SET
	DEFAULT now(),
ALTER COLUMN
	"updated_at"
SET
	DEFAULT now();

-- Modify "company_addresses" table
ALTER TABLE
	"public"."company_addresses"
ALTER COLUMN
	"created_at"
SET
	DEFAULT now(),
ALTER COLUMN
	"updated_at"
SET
	DEFAULT now();

-- Modify "company_categories" table
ALTER TABLE
	"public"."company_categories"
ALTER COLUMN
	"created_at"
SET
	DEFAULT now(),
ALTER COLUMN
	"updated_at"
SET
	DEFAULT now();

-- Modify "states" table
ALTER TABLE
	"public"."states"
ALTER COLUMN
	"created_at"
SET
	DEFAULT now();