\ir ./setup-db-primary.sql;

CREATE OR REPLACE VIEW "ViewSum3" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1ID",
    "Table2"."ID" AS "Table2ID",
    "Table1"."I" AS "Table1I",
    "Table2"."I" AS "Table2I",
    "Table1"."Num1" AS "Table1Num1",
    "Table1"."Num2" AS "Table1Num2",
    "Table2"."Num3" AS "Table2Num3",
    "Table2"."Num4" AS "Table2Num4",
    "Table1"."Num2" + "Table1"."Num2" + "Table2"."Num3" + "Table2"."Num4" AS "Sum1234"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum3()
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
        TRUE AS view,
        current_database() AS db
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
        'Table2' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS view,
        current_database() AS db
      FROM (
        SELECT
          NEW."Table2ID" AS "ID",
          NEW."Table2I" AS "I"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Table1Num1" <> OLD."Table1Num1") THEN
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
          NEW."Table1Num1" AS "Num1"
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
  IF (NEW."Table2Num3" <> OLD."Table2Num3") THEN
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
          NEW."Table2Num3" AS "Num3"
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

DROP TRIGGER IF EXISTS "ViewSum3_process" ON "ViewSum3";
CREATE TRIGGER "ViewSum3_process"
INSTEAD OF INSERT OR UPDATE OR DELETE ON "ViewSum3"
FOR EACH ROW
EXECUTE PROCEDURE process_viewsum3();

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum3_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
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

CREATE OR REPLACE FUNCTION notify_from_table1_viewsum3_after()
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
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
        WHERE "Table1ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum3_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewsum3_after()
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
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
        WHERE "Table2ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewSum3' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewSum3"
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

DROP TRIGGER IF EXISTS "Table1_notify_viewsum3_before" ON "Table1";
CREATE TRIGGER "Table1_notify_viewsum3_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum3_before();

DROP TRIGGER IF EXISTS "Table1_notify_viewsum3_after" ON "Table1";
CREATE TRIGGER "Table1_notify_viewsum3_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewsum3_after();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum3_before" ON "Table2";
CREATE TRIGGER "Table2_notify_viewsum3_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum3_before();

DROP TRIGGER IF EXISTS "Table2_notify_viewsum3_after" ON "Table2";
CREATE TRIGGER "Table2_notify_viewsum3_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewsum3_after();

DROP TRIGGER IF EXISTS "Table1_notify" ON "Table1";
DROP TRIGGER IF EXISTS "Table2_notify" ON "Table2";

/*
  Delete the databases before running this function.
  The IDs MUST be equal everywhere
*/
CREATE OR REPLACE FUNCTION goahead(i BIGINT)
  RETURNS VOID
  LANGUAGE 'plpgsql'
AS $$
DECLARE
  c111 double precision=CEIL(RANDOM() * 1000) / 10;
  c112 double precision=CEIL(RANDOM() * 1000) / 10;
  c123 double precision=CEIL(RANDOM() * 1000) / 10;
  c124 double precision=CEIL(RANDOM() * 1000) / 10;
  c211 double precision=CEIL(RANDOM() * 1000) / 10;
  c212 double precision=CEIL(RANDOM() * 1000) / 10;
  c223 double precision=CEIL(RANDOM() * 1000) / 10;
  c224 double precision=CEIL(RANDOM() * 1000) / 10;
BEGIN
RAISE NOTICE 'c111=% c112=% c123=% c124=%', c111, c112, c123, c124;
RAISE NOTICE 'c211=% c212=% c223=% c224=%', c211, c212, c223, c224;
UPDATE "ViewSum3"
  SET
    "Table1Num1"=c111,
    "Table1Num2"=c112,
    "Table2Num3"=c123,
    "Table2Num4"=c124,
    "Table1I"=i,
    "Table2I"=i
  WHERE "ID"='1:1';
UPDATE "ViewSum3"
  SET
    "Table1Num1"=c211,
    "Table1Num2"=c212,
    "Table2Num3"=c223,
    "Table2Num4"=c224,
    "Table1I"=i,
    "Table2I"=i
  WHERE "ID"='2:2';
END;
$$;