
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
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1Num2" <> OLD."Table1Num2") THEN
    UPDATE "Table1"
      SET "Num2"=NEW."Table1Num2"
      WHERE "ID"=NEW."Table1ID";
  END IF;
  IF (NEW."Table2Num4" <> OLD."Table2Num4") THEN
    UPDATE "Table2"
      SET "Num4"=NEW."Table2Num4"
      WHERE "ID"=NEW."Table2ID";
  END IF;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;