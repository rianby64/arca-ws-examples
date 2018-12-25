
\ir ./setup-db1.sql;

DROP VIEW IF EXISTS "viewSum1" CASCADE;

CREATE OR REPLACE VIEW "viewSum1" AS (
  SELECT
    "Table1"."ID" || ':' || "Table2"."ID" AS "ID",
    "Table1"."ID" AS "Table1_ID",
    "Table2"."ID" AS "Table2_ID",
    "Table1"."Num1" AS "Table1_Num1",
    "Table2"."Num3" AS "Table2_Num3",
    "Table1"."Num1" + "Table2"."Num3" AS "Sum13"
  FROM "Table1", "Table2"
);

CREATE OR REPLACE FUNCTION process_viewsum1()
  RETURNS trigger
  LANGUAGE 'plpgsql'
  IMMUTABLE
AS $$
BEGIN
IF TG_OP = 'UPDATE' THEN
  IF (NEW."Table1_Num1" <> OLD."Table1_Num1") THEN
    UPDATE "Table1"
      SET "Table1_Num1"=NEW."Table1_Num1"
      WHERE "ID"=NEW."Table1_ID";
  END IF;
  IF (NEW."Table2_Num3" <> OLD."Table2_Num3") THEN
    UPDATE "Table2"
      SET "Table2_Num3"=NEW."Table2_Num3"
      WHERE "ID"=NEW."Table2_ID";
  END IF;
  RETURN NEW;
END IF;
RETURN NULL;
END;
$$;