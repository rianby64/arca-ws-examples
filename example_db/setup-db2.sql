
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