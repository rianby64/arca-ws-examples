
DROP FUNCTION IF EXISTS notify_jsonrpc CASCADE;
CREATE FUNCTION notify_jsonrpc()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    IMMUTABLE
AS $$
BEGIN
IF (TG_OP = 'INSERT') THEN
  PERFORM pg_notify('jsonrpc', json_build_object(
    'source', TG_TABLE_NAME,
    'method', LOWER(TG_OP),
    'result', row_to_json(NEW))::text);
ELSIF (TG_OP = 'DELETE') THEN
  PERFORM pg_notify('jsonrpc', json_build_object(
    'source', TG_TABLE_NAME,
    'method', LOWER(TG_OP),
    'result', row_to_json(OLD))::text);
ELSIF (TG_OP = 'UPDATE') THEN
  PERFORM pg_notify('jsonrpc', json_build_object(
    'source', TG_TABLE_NAME,
    'method', LOWER(TG_OP),
    'result', row_to_json(NEW))::text);
END IF;
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