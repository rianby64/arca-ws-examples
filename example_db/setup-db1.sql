
/*
  Normaliza el keynote
  Tenga presente que aqui puede ir cualquier cantidad
  de separadores. Actualmente está configurado para
  delimintar hasta 8 grupos. En caso de necesitar más
  grupos basta con replicar simetricamente el código.
*/
CREATE OR REPLACE FUNCTION normalize_keynote(keynote character varying(255))
  RETURNS character varying(255) AS
$BODY$
BEGIN
  RETURN
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 1))) ||
    SPLIT_PART(keynote, '.', 1) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 2))) ||
    SPLIT_PART(keynote, '.', 2) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 3))) ||
    SPLIT_PART(keynote, '.', 3) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 4))) ||
    SPLIT_PART(keynote, '.', 4) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 5))) ||
    SPLIT_PART(keynote, '.', 5) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 6))) ||
    SPLIT_PART(keynote, '.', 6) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 7))) ||
    SPLIT_PART(keynote, '.', 7) || '.' ||
    REPEAT('0', 5 - LENGTH(SPLIT_PART(keynote, '.', 8))) ||
    SPLIT_PART(keynote, '.', 8);
END;
$BODY$
LANGUAGE plpgsql IMMUTABLE;

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

CREATE OR REPLACE FUNCTION a_setup_from_aau_dui_before()
  RETURNS TRIGGER AS
$$
BEGIN
  IF (TG_OP = 'INSERT') THEN
    NEW."Expand" = FALSE;
    IF (NEW."Parent" IS NULL) THEN
      NEW."Expand" = TRUE;
      RETURN NEW;
    END IF;
    IF (NEW."ID" IS NOT NULL AND NEW."Parent" IS NOT NULL) THEN
      RETURN NEW;
    END IF;
    NEW."ID" = (SELECT NEW."Parent" || '.' || (COUNT('id') + 1)::text
      FROM "AAU"
      WHERE "Parent" LIKE NEW."Parent");
    RETURN NEW;
  ELSIF (TG_OP = 'DELETE') THEN
    IF (substring(OLD."ID", 1, 1) = '-') THEN
      RETURN NULL;
    END IF;
    RETURN OLD;
  ELSIF (TG_OP = 'UPDATE') THEN
    IF (substring(OLD."ID", 1, 1) = '-') THEN
      IF (OLD."ID" != NEW."ID") THEN
        RAISE NOTICE 'Cannot change AAU."ID"=%', OLD."ID";
        RETURN NULL;
      END IF;
    END IF;
    IF (substring(OLD."Parent", 1, 1) = '-') THEN
      IF (OLD."Parent" != NEW."Parent") THEN
        RAISE NOTICE 'Cannot change AAU."Parent"=%', OLD."Parent";
        RETURN NULL;
      END IF;
    END IF;
    NEW."Expand" = (SELECT COUNT("ID") > 0
      FROM "AAU"
      WHERE "Parent" LIKE NEW."ID");
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION a_setup_from_aau_dui_after()
  RETURNS TRIGGER AS
$$
DECLARE
  deconcretized_id character varying(255);
  deconcretized_parent character varying(255);
BEGIN
  IF (TG_OP = 'INSERT') THEN
    IF (substring(NEW."ID", 1, 1) <> '-' AND NEW."Parent" IS NOT NULL) THEN
      deconcretized_id := regexp_replace(NEW."ID", E'^[0-9]+[.]', '-.');
      deconcretized_parent := regexp_replace(NEW."Parent", E'^[0-9]+[.]', '-.');
      IF (deconcretized_parent = NEW."Parent" AND NEW."Parent" IS NOT NULL) THEN
        deconcretized_parent = '-';
      END IF;
      IF (deconcretized_id <> NEW."ID") THEN
        IF ((SELECT "ID"
          FROM "AAU"
          WHERE "ID"=deconcretized_id AND "Parent"=deconcretized_parent
          LIMIT 1) IS NULL) THEN
          INSERT INTO "AAU" ("ID", "Parent")
            VALUES (deconcretized_id, deconcretized_parent);
        END IF;
      END IF;
    END IF;
    UPDATE "AAU"
      SET "Expand"=NULL
      WHERE "ID"=NEW."Parent" AND "Parent" IS NOT NULL;
    RETURN NEW;
  ELSIF (TG_OP = 'DELETE') THEN
    IF (OLD."Parent" IS NULL) THEN
      DELETE FROM "Projects" WHERE "ID"=OLD."ID"::integer;
    END IF;
    IF (OLD."Parent" IS NOT NULL) THEN
      UPDATE "AAU"
        SET "Expand"=NULL
        WHERE "ID"=OLD."Parent" AND "Parent" IS NOT NULL;
    END IF;
    RETURN OLD;
  ELSIF (TG_OP = 'UPDATE') THEN
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS "AAU"
(
  "ID" character varying(255) NOT NULL,
  "Parent" character varying(255),
  "Expand" boolean NOT NULL DEFAULT FALSE,
  "Description" character varying(255) NOT NULL DEFAULT '',
  "Information" TEXT,
  "Unit" character varying(31),
  "Qop" double precision,
  "Estimated" double precision,
  "CreatedAt" timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT "AAU_pkey" PRIMARY KEY ("ID"),
  CONSTRAINT "AAU_parent_fkey" FOREIGN KEY ("Parent")
      REFERENCES "AAU" ("ID") MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "AAU_id_unique" UNIQUE ("ID")
);

DROP TRIGGER IF EXISTS "AAU_notify" ON "AAU" CASCADE;
CREATE TRIGGER "AAU_notify"
  AFTER INSERT OR UPDATE OR DELETE
  ON "AAU"
  FOR EACH ROW
  EXECUTE PROCEDURE notify_jsonrpc();

DROP TRIGGER IF EXISTS "AAU_a_setup_DUI_before" ON "AAU" CASCADE;
CREATE TRIGGER "AAU_a_setup_DUI_before"
  BEFORE INSERT OR UPDATE OR DELETE
  ON "AAU"
  FOR EACH ROW
  EXECUTE PROCEDURE a_setup_from_aau_dui_before();

DROP TRIGGER IF EXISTS "AAU_a_setup_DUI_after" ON "AAU" CASCADE;
CREATE TRIGGER "AAU_a_setup_DUI_after"
  AFTER INSERT OR UPDATE OR DELETE
  ON "AAU"
  FOR EACH ROW
  EXECUTE PROCEDURE a_setup_from_aau_dui_after();

DO
$BODY$
BEGIN
  IF ((SELECT "ID" FROM "AAU" WHERE "ID"='-') IS NULL) THEN
    INSERT INTO "AAU" ("ID", "Parent", "Expand") VALUES ('-', NULL, TRUE);
  END IF;
END;
$BODY$;