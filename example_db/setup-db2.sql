
\ir ./setup-db1.sql;

DROP VIEW IF EXISTS "ViewSum1" CASCADE;

CREATE OR REPLACE VIEW "ViewSum1" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1ID",
    "Table2"."ID" AS "Table2ID",
    "Table1"."Num1" AS "Table1Num1",
    "Table2"."Num3" AS "Table2Num3",
    "Table1"."Num1" + "Table2"."Num3" AS "Sum13"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum1()
  RETURNS trigger
  LANGUAGE 'plpgsql'
  VOLATILE
AS $$
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1Num1" <> OLD."Table1Num1") THEN
    UPDATE "Table1"
      SET "Num1"=NEW."Table1Num1"
      WHERE "ID"=NEW."Table1ID";
  END IF;
  IF (NEW."Table2Num3" <> OLD."Table2Num3") THEN
    UPDATE "Table2"
      SET "Num3"=NEW."Table2Num3"
      WHERE "ID"=NEW."Table2ID";
  END IF;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS "ViewSum1_process" ON "ViewSum1" CASCADE;
CREATE TRIGGER "ViewSum1_process"
INSTEAD OF INSERT OR UPDATE OR DELETE ON "ViewSum1"
FOR EACH ROW
EXECUTE PROCEDURE process_viewsum1();

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum1_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
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

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum1_after()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
      WHERE "Table1ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum1_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum1_after()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  RETURN OLD;
ELSIF (TG_OP = 'UPDATE') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
      WHERE "Table2ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (SELECT
      'ViewSum1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (SELECT "ViewSum1".*
      FROM "ViewSum1"
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

DROP TRIGGER IF EXISTS "Table1_notify_viewsum1_before" ON "Table1" CASCADE;
CREATE TRIGGER "Table1_notify_viewsum1_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum1_before();

DROP TRIGGER IF EXISTS "Table1_notify_viewsum1_after" ON "Table1" CASCADE;
CREATE TRIGGER "Table1_notify_viewsum1_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum1_after();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum1_before" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify_viewsum1_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum1_before();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum1_after" ON "Table2" CASCADE;
CREATE TRIGGER "Table2_notify_viewsum1_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum1_after();