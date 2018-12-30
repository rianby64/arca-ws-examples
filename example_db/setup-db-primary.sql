\set QUIET 1

CREATE TABLE IF NOT EXISTS "Table1"
(
  "ID" SERIAL,
  "I" bigint DEFAULT 0,
  "Num1" double precision DEFAULT 0.0,
  "Num2" double precision DEFAULT 0.0,
  "CreatedAt" timestamp with time zone DEFAULT now(),
  CONSTRAINT "Table1_pkey" PRIMARY KEY ("ID")
);

TRUNCATE "Table1";
INSERT INTO "Table1"("Num1", "Num2")
  VALUES (1.0, 2.0), (3.0, 4.0);

CREATE TABLE IF NOT EXISTS "Table2"
(
  "ID" SERIAL,
  "I" bigint DEFAULT 0,
  "Num3" double precision DEFAULT 0.0,
  "Num4" double precision DEFAULT 0.0,
  "CreatedAt" timestamp with time zone DEFAULT now(),
  CONSTRAINT "Table2_pkey" PRIMARY KEY ("ID")
);

TRUNCATE "Table2";
INSERT INTO "Table2"("Num3", "Num4")
  VALUES (5.0, 6.0), (7.0, 8.0);