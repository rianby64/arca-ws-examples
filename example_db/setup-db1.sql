
DROP FUNCTION IF EXISTS send_row_jsonrpc CASCADE;
CREATE FUNCTION send_row_jsonrpc(source TEXT, method TEXT, row JSON)
    RETURNS VOID
    LANGUAGE 'plpgsql'
    IMMUTABLE
AS $$
BEGIN
PERFORM pg_notify('jsonrpc'::text, json_build_object(
  'source', source,
  'method', LOWER(method),
  'result', row)::text);
END;
$$;

DROP FUNCTION IF EXISTS notify_jsonrpc CASCADE;
CREATE FUNCTION notify_jsonrpc()
    RETURNS TRIGGER
    LANGUAGE 'plpgsql'
    IMMUTABLE
AS $$
DECLARE
  rec RECORD;
BEGIN
IF (TG_OP = 'INSERT') THEN
  rec := NEW;
ELSIF (TG_OP = 'DELETE') THEN
  rec := OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  rec := NEW;
END IF;
PERFORM send_row_jsonrpc(TG_TABLE_NAME, TG_OP, row_to_json(REC));
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

DROP TRIGGER IF EXISTS "Table2_notify" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_jsonrpc();