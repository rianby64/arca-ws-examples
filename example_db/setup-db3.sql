
\ir ./setup-db1.sql;

DROP VIEW IF EXISTS "viewSum2" CASCADE;

CREATE OR REPLACE VIEW "viewSum2" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1_ID",
    "Table2"."ID" AS "Table2_ID",
    "Table1"."Num2" AS "Table1_Num2",
    "Table2"."Num4" AS "Table2_Num4",
    "Table1"."Num2" + "Table2"."Num4" AS "Sum24"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum2()
  RETURNS trigger
  LANGUAGE 'plpgsql'
  IMMUTABLE
AS $$
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1_Num2" <> OLD."Table1_Num2") THEN
    UPDATE "Table1"
      SET "Table1_Num2"=NEW."Table1_Num2"
      WHERE "ID"=NEW."Table1_ID";
  END IF;
  IF (NEW."Table2_Num4" <> OLD."Table2_Num4") THEN
    UPDATE "Table2"
      SET "Table2_Num4"=NEW."Table2_Num4"
      WHERE "ID"=NEW."Table2_ID";
  END IF;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;