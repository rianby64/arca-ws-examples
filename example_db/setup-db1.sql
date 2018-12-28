\set QUIET 1

DROP FUNCTION IF EXISTS notify_jsonrpc CASCADE;
CREATE OR REPLACE FUNCTION notify_jsonrpc()
    RETURNS TRIGGER
    LANGUAGE 'plpgsql'
    IMMUTABLE
AS $$
DECLARE
  rec RECORD;
  cmp RECORD;
  msg TEXT;
  tst TEXT;
BEGIN
IF (TG_OP = 'INSERT') THEN
  rec := NEW;
  cmp := OLD;
ELSIF (TG_OP = 'DELETE') THEN
  rec := OLD;
  cmp := NEW;
ELSIF (TG_OP = 'UPDATE') THEN
  rec := NEW;
  cmp := OLD;
END IF;
msg := json_build_object(
  'source', TG_TABLE_NAME,
  'method', lower(TG_OP),
  'db', current_database(),
  'primary', TRUE,
  'result', row_to_json(rec))::text;
tst := json_build_object(
  'source', TG_TABLE_NAME,
  'method', lower(TG_OP),
  'db', current_database(),
  'primary', TRUE,
  'result', row_to_json(cmp))::text;
IF (tst = msg) THEN
  RETURN NULL;
END IF;
PERFORM pg_notify('jsonrpc', msg);
RETURN NULL;
END;
$$;

CREATE TABLE IF NOT EXISTS "Table1"
(
  "ID" SERIAL,
  "Num1" double precision DEFAULT 0.0,
  "Num2" double precision DEFAULT 0.0,
  "CreatedAt" timestamp with time zone DEFAULT now(),
  CONSTRAINT "Table1_pkey" PRIMARY KEY ("ID")
);

TRUNCATE "Table1";
INSERT INTO "Table1"("Num1", "Num2")
  VALUES (1.0, 2.0), (3.0, 4.0);

DROP TRIGGER IF EXISTS "Table1_notify" ON "Table1" CASCADE;
CREATE TRIGGER "Table1_notify"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_jsonrpc();


CREATE TABLE IF NOT EXISTS "Table2"
(
  "ID" SERIAL,
  "Num3" double precision DEFAULT 0.0,
  "Num4" double precision DEFAULT 0.0,
  "CreatedAt" timestamp with time zone DEFAULT now(),
  CONSTRAINT "Table2_pkey" PRIMARY KEY ("ID")
);

TRUNCATE "Table2";
INSERT INTO "Table2"("Num3", "Num4")
  VALUES (5.0, 6.0), (7.0, 8.0);

DROP TRIGGER IF EXISTS "Table2_notify" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_jsonrpc();