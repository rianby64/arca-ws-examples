\ir ./setup-db-primary.sql;

CREATE OR REPLACE VIEW "ViewTable1" AS (
  SELECT
    "Table1"."ID" AS "ID",
    "Table1"."I" AS "I",
    "Table1"."Num1" AS "Num1",
    "Table1"."Num2" AS "Num2"
  FROM "Table1"
);

CREATE OR REPLACE FUNCTION process_viewtable1()
  RETURNS trigger
  LANGUAGE 'plpgsql'
AS $$
DECLARE
  r RECORD;
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."I" <> OLD."I") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."I" AS "I"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Num1" <> OLD."Num1") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."Num1" AS "Num1"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Num2" <> OLD."Num2") THEN
    FOR r IN (
      SELECT
        'Table1' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."Num2" AS "Num2"
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

CREATE OR REPLACE FUNCTION notify_from_table1_viewtable1_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewTable1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable1"
        WHERE "ID"=OLD."ID"
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

CREATE OR REPLACE FUNCTION notify_from_table1_viewtable1_after()
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
      'ViewTable1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable1"
        WHERE "ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewTable1' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable1"
        WHERE "ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS "ViewTable1_process" ON "ViewTable1";
CREATE TRIGGER "ViewTable1_process"
INSTEAD OF INSERT OR UPDATE OR DELETE ON "ViewTable1"
FOR EACH ROW
EXECUTE PROCEDURE process_viewtable1();

DROP TRIGGER IF EXISTS "Table1_notify_viewTable1_before" ON "Table1";
CREATE TRIGGER "Table1_notify_viewTable1_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewtable1_before();

DROP TRIGGER IF EXISTS "Table1_notify_viewTable1_after" ON "Table1";
CREATE TRIGGER "Table1_notify_viewTable1_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table1"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table1_viewtable1_after();

CREATE OR REPLACE VIEW "ViewTable2" AS (
  SELECT
    "Table2"."ID" AS "ID",
    "Table2"."I" AS "I",
    "Table2"."Num3" AS "Num3",
    "Table2"."Num4" AS "Num4"
  FROM "Table2"
);

CREATE OR REPLACE FUNCTION process_viewtable2()
  RETURNS trigger
  LANGUAGE 'plpgsql'
AS $$
DECLARE
  r RECORD;
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."I" <> OLD."I") THEN
    FOR r IN (
      SELECT
        'Table2' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."I" AS "I"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Num3" <> OLD."Num3") THEN
    FOR r IN (
      SELECT
        'Table2' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."Num3" AS "Num3"
      ) t
    ) LOOP
      PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
    END LOOP;
  END IF;
  IF (NEW."Num4" <> OLD."Num4") THEN
    FOR r IN (
      SELECT
        'Table2' AS source,
        lower(TG_OP) AS method,
        row_to_json(t) AS result,
        TRUE AS primary,
        current_database() AS db
      FROM (
        SELECT
          NEW."ID" AS "ID",
          NEW."Num4" AS "Num4"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewtable2_before()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
DECLARE
  r RECORD;
BEGIN
IF (TG_OP = 'DELETE') THEN
  FOR r IN (
    SELECT
      'ViewTable2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable2"
        WHERE "ID"=OLD."ID"
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

CREATE OR REPLACE FUNCTION notify_from_table2_viewtable2_after()
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
      'ViewTable2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable2"
        WHERE "ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
ELSIF (TG_OP = 'INSERT') THEN
  FOR r IN (
    SELECT
      'ViewTable2' AS source,
      lower(TG_OP) AS method,
      row_to_json(t) AS result,
      current_database() AS db
    FROM (
      SELECT *
        FROM "ViewTable2"
        WHERE "ID"=NEW."ID"
    ) t
  ) LOOP
    PERFORM pg_notify('jsonrpc', row_to_json(r)::text);
  END LOOP;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS "ViewTable2_process" ON "ViewTable2";
CREATE TRIGGER "ViewTable2_process"
INSTEAD OF INSERT OR UPDATE OR DELETE ON "ViewTable2"
FOR EACH ROW
EXECUTE PROCEDURE process_viewtable2();

DROP TRIGGER IF EXISTS "Table2_notify_viewTable2_before" ON "Table2";
CREATE TRIGGER "Table2_notify_viewTable2_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewtable2_before();

DROP TRIGGER IF EXISTS "Table2_notify_viewTable2_after" ON "Table2";
CREATE TRIGGER "Table2_notify_viewTable2_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "Table2"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_from_table2_viewtable2_after();