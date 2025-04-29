-- cria ou atualiza o role de réplica para todo o cluster
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='replicationuser') THEN
    CREATE ROLE replicationUser
      WITH REPLICATION LOGIN
      PASSWORD 'admin';
  ELSE
    ALTER ROLE replicationUser
      WITH ENCRYPTED PASSWORD 'admin';
  END IF;
END
$$;

------------------------------------------------------------------
--- Level 1 -> Product
------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS "customers" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "email" VARCHAR(255) DEFAULT NULL,
  "password" VARCHAR(255) DEFAULT NULL,
  "first_name" VARCHAR(255) DEFAULT NULL,
  "last_name" VARCHAR(255) DEFAULT NULL,
  "cpf" VARCHAR(255) DEFAULT NULL,
  "date_of_birth" TIMESTAMP DEFAULT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "addresses" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "customer_id" bigint DEFAULT NULL,
  "street" VARCHAR(255) DEFAULT NULL,
  "number" VARCHAR(255) DEFAULT NULL,
  "complement" VARCHAR(255) DEFAULT NULL,
  "city" VARCHAR(255) DEFAULT NULL,
  "state" VARCHAR(255) DEFAULT NULL,
  "postal_code" VARCHAR(255) DEFAULT NULL,
  "country" VARCHAR(255) DEFAULT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

------------------------------------------------------------------
-- INDEX
------------------------------------------------------------------

-- CUSTOMER
CREATE INDEX IF NOT EXISTS "idx_customers_id" ON "customers" ("id");
CREATE INDEX IF NOT EXISTS "idx_customers_name" ON "customers" ("first_name", "last_name");
CREATE INDEX IF NOT EXISTS "idx_customers_cpf" ON "customers" ("cpf");
CREATE INDEX IF NOT EXISTS "idx_customers_email" ON "customers" ("email");
CREATE INDEX IF NOT EXISTS "idx_customers_date_of_birth" ON "customers" ("date_of_birth");

-- ADDRESS
CREATE INDEX IF NOT EXISTS "idx_addresses_customer_id" ON "addresses" ("customer_id");
CREATE INDEX IF NOT EXISTS "idx_addresses_id" ON "addresses" ("id");
CREATE INDEX IF NOT EXISTS "idx_addresses_street" ON "addresses" ("street");
CREATE INDEX IF NOT EXISTS "idx_addresses_city" ON "addresses" ("city");
CREATE INDEX IF NOT EXISTS "idx_addresses_state" ON "addresses" ("state");
CREATE INDEX IF NOT EXISTS "idx_addresses_country" ON "addresses" ("country");


------------------------------------------------------------------
-- CONSTRAINTS
------------------------------------------------------------------
ALTER TABLE "addresses" ADD CONSTRAINT "fk_customer_id" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");


--
--
-- ------------------------------------------------------------------
-- -- LIMPEZA DA BASE
-- ------------------------------------------------------------------
--

TRUNCATE TABLE addresses RESTART IDENTITY CASCADE;
TRUNCATE TABLE customers RESTART IDENTITY CASCADE;
