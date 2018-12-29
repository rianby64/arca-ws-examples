
\ir ./setup-db1.sql;

DROP VIEW IF EXISTS "ViewSum2" CASCADE;

CREATE OR REPLACE VIEW "ViewSum2" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1ID",
    "Table2"."ID" AS "Table2ID",
    "Table1"."Num2" AS "Table1Num2",
    "Table2"."Num4" AS "Table2Num4",
    "Table1"."Num2" + "Table2"."Num4" AS "Sum24"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum2()
  RETURNS trigger
  LANGUAGE 'plpgsql'
  VOLATILE
AS $$
DECLARE
  r RECORD;
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1Num2" <> OLD."Table1Num2") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS view,
        current_database() AS db
      FROM (
        SELECT
          NEW."Table1ID" AS "ID",
          NEW."Table1Num2" AS "Num2"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Table2Num4" <> OLD."Table2Num4") THEN
    FOR r IN (
      SELECT
        'Table2' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS view,
        current_database() AS db
      FROM (
        SELECT
          NEW."Table2ID" AS "ID",
          NEW."Table2Num4" AS "Num4"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS "ViewSum2_process" ON "ViewSum2" CASCADE;
CREATE TRIGGER "ViewSum2_process"
INSTEAD OF INSERT OR UPDATE OR DELETE ON "ViewSum2"
FOR EACH ROW
EXECUTE PROCEDURE process_viewsum2();

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum2_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table1ID"=OLD."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum2_after()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table1ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table1ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum2_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table2ID"=OLD."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum2_after()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table2ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewSum2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum2"
        WHERE "Table2ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS "Table1_notify_viewsum2_before" ON "Table1" CASCADE;
CREATE TRIGGER "Table1_notify_viewsum2_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum2_before();

DROP TRIGGER IF EXISTS "Table1_notify_viewsum2_after" ON "Table1" CASCADE;
CREATE TRIGGER "Table1_notify_viewsum2_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum2_after();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum2_before" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify_viewsum2_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum2_before();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum2_after" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify_viewsum2_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum2_after();

DROP TRIGGER IF EXISTS "Table1_notify" ON "Table1" CASCADE;
DROP TRIGGER IF EXISTS "Table2_notify" ON "Table2" CASCADE;