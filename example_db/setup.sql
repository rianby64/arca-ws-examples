
DROP FUNCTION IF EXISTS notify_jsonrpc CASCADE;
CREATE FUNCTION notify_jsonrpc()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    IMMUTABLE
AS $BODY$
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
$BODY$;

DROP TRIGGER IF EXISTS "test_notify" ON "test" CASCADE;
CREATE TRIGGER "test_notify"
  AFTER INSERT OR UPDATE OR DELETE
  ON "test"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_jsonrpc();