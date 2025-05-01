-- ==============================================================
-- 1. Create replication role
-- ==============================================================

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='replicationuser') THEN
    CREATE ROLE replicationuser
      WITH REPLICATION LOGIN
      PASSWORD 'admin';
  ELSE
    ALTER ROLE replicationuser
      WITH ENCRYPTED PASSWORD 'admin';
  END IF;
END
$$;

-- ==============================================================
-- 2. Create SEQUENCES (one per table)
-- ==============================================================

CREATE SEQUENCE IF NOT EXISTS id_sequence_customers;
CREATE SEQUENCE IF NOT EXISTS id_sequence_addresses;

-- ==============================================================
-- 3. Instagram-style ID generation function (with clock and sequence control)
-- ==============================================================

CREATE OR REPLACE FUNCTION generate_instagram_id(
  shard_id INTEGER,
  sequence_name TEXT
)
RETURNS BIGINT AS $$
DECLARE
    our_epoch      BIGINT := 1609459200000;  -- 2021-01-01 00:00:00 UTC
    timestamp_ms   BIGINT;
    last_timestamp BIGINT;
    now_ms         BIGINT;
    raw_seq        BIGINT;
    sequence_id    BIGINT;
    elapsed        BIGINT;
    result         BIGINT;
BEGIN
    IF shard_id < 0 OR shard_id > 8191 THEN
        RAISE EXCEPTION 'shard_id out of allowed range (0–8191)';
    END IF;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO timestamp_ms;
    last_timestamp := timestamp_ms;

    SELECT FLOOR(EXTRACT(EPOCH FROM now()) * 1000) INTO now_ms;
    IF timestamp_ms > now_ms + 5000 THEN
        RAISE EXCEPTION 'System clock appears ahead: timestamp_ms=%, now_ms=%', timestamp_ms, now_ms;
    END IF;

    LOOP
        EXECUTE format('SELECT nextval(%L)', sequence_name) INTO raw_seq;
        sequence_id := raw_seq % 1024;

        IF sequence_id < 1023 THEN
            EXIT;
        END IF;

        PERFORM pg_sleep(0.001);
        SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO timestamp_ms;

        IF timestamp_ms > last_timestamp THEN
              EXECUTE format('SELECT setval(%L, 0, false)', sequence_name);
            last_timestamp := timestamp_ms;
        END IF;
    END LOOP;

    elapsed := timestamp_ms - our_epoch;
    result := (elapsed << 23) | (shard_id << 10) | sequence_id;
    RETURN result;
END;
$$ LANGUAGE plpgsql;

-- ==============================================================
-- 4. Tables using their respective sequences in ID
-- ==============================================================

CREATE TABLE IF NOT EXISTS "customers" (
  "id" BIGINT PRIMARY KEY DEFAULT generate_instagram_id(1, 'id_sequence_customers'),
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
  "id" BIGINT PRIMARY KEY DEFAULT generate_instagram_id(1, 'id_sequence_addresses'),
    "customer_id" BIGINT DEFAULT NULL,
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

-- ==============================================================
-- 5. Indexes and constraints
-- ==============================================================

-- Indexes for CUSTOMERS
CREATE INDEX IF NOT EXISTS "idx_customers_id" ON "customers" ("id");
CREATE INDEX IF NOT EXISTS "idx_customers_name" ON "customers" ("first_name", "last_name");
CREATE INDEX IF NOT EXISTS "idx_customers_cpf" ON "customers" ("cpf");
CREATE INDEX IF NOT EXISTS "idx_customers_email" ON "customers" ("email");
CREATE INDEX IF NOT EXISTS "idx_customers_date_of_birth" ON "customers" ("date_of_birth");

-- Indexes for ADDRESSES
CREATE INDEX IF NOT EXISTS "idx_addresses_customer_id" ON "addresses" ("customer_id");
CREATE INDEX IF NOT EXISTS "idx_addresses_id" ON "addresses" ("id");
CREATE INDEX IF NOT EXISTS "idx_addresses_street" ON "addresses" ("street");
CREATE INDEX IF NOT EXISTS "idx_addresses_city" ON "addresses" ("city");
CREATE INDEX IF NOT EXISTS "idx_addresses_state" ON "addresses" ("state");
CREATE INDEX IF NOT EXISTS "idx_addresses_country" ON "addresses" ("country");

-- Foreign key
ALTER TABLE "addresses" ADD CONSTRAINT "fk_customer_id" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

-- ==============================================================
-- 6. Truncation (no RESTART IDENTITY since IDs are custom generated)
-- ==============================================================

TRUNCATE TABLE addresses CASCADE;
TRUNCATE TABLE customers CASCADE;

