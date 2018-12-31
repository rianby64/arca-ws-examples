\ir ./setup-db-primary.sql;

CREATE OR REPLACE VIEW "ViewSum2" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1ID",
    "Table2"."ID" AS "Table2ID",
    "Table1"."I" AS "Table1I",
    "Table2"."I" AS "Table2I",
    "Table1"."Num2" AS "Table1Num2",
    "Table2"."Num4" AS "Table2Num4",
    "Table1"."Num2" + "Table2"."Num4" AS "Sum24"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum2()
  RETURNS trigger
  LANGUAGE 'plpgsql'
AS $$
DECLARE
  r RECORD;
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1I" <> OLD."Table1I") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary
      FROM (
        SELECT
          NEW."Table1ID" AS "ID",
          NEW."Table1I" AS "I"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Table2I" <> OLD."Table2I") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary
      FROM (
        SELECT
          NEW."Table2ID" AS "ID",
          NEW."Table2I" AS "I"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Table1Num2" <> OLD."Table1Num2") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary
      FROM (
        SELECT
          NEW."Table1ID" AS "ID",
          NEW."Table1I" AS "I",
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
        TRUE AS primary
      FROM (
        SELECT
          NEW."Table2ID" AS "ID",
          NEW."Table2I" AS "I",
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

DROP TRIGGER IF EXISTS "ViewSum2_process" ON "ViewSum2";
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
      row_to_json(t) AS result
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
      row_to_json(t) AS result
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
      row_to_json(t) AS result
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
      row_to_json(t) AS result
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
      row_to_json(t) AS result
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
      row_to_json(t) AS result
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

DROP TRIGGER IF EXISTS "Table1_notify_viewsum2_before" ON "Table1";
CREATE TRIGGER "Table1_notify_viewsum2_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum2_before();

DROP TRIGGER IF EXISTS "Table1_notify_viewsum2_after" ON "Table1";
CREATE TRIGGER "Table1_notify_viewsum2_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum2_after();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum2_before" ON "Table2";
CREATE TRIGGER "Table2_notify_viewsum2_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum2_before();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum2_after" ON "Table2";
CREATE TRIGGER "Table2_notify_viewsum2_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum2_after();

/*
  Delete the databases before running this function.
  The IDs MUST be equal everywhere
*/
CREATE OR REPLACE FUNCTION goahead(i BIGINT)
  RETURNS VOID
  LANGUAGE 'plpgsql'
AS $$
DECLARE
  c112 double precision=CEIL(RANDOM() * 1000) / 10;
  c124 double precision=CEIL(RANDOM() * 1000) / 10;
  c212 double precision=CEIL(RANDOM() * 1000) / 10;
  c224 double precision=CEIL(RANDOM() * 1000) / 10;
BEGIN
RAISE NOTICE 'I=% c112=% c124=%', i, c112, c124;
RAISE NOTICE 'I=% c212=% c224=%', i, c212, c224;
UPDATE "ViewSum2"
  SET
    "Table1Num2"=c112,
    "Table2Num4"=c124,
    "Table1I"=i,
    "Table2I"=i
  WHERE "ID"='1:1';
UPDATE "ViewSum2"
  SET
    "Table1Num2"=c212,
    "Table2Num4"=c224,
    "Table1I"=i,
    "Table2I"=i
  WHERE "ID"='2:2';
END;
$$;