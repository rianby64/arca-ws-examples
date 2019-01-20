
\unset ECHO
\set QUIET 1
-- Turn off echo and keep things quiet.

-- Format the output for nice TAP.
\pset format unaligned
\pset tuples_only true
\pset pager off

-- Revert all changes on failure.
\set ON_ERROR_ROLLBACK 1
\set ON_ERROR_STOP true

\ir ./main.sql
\ir ./process_jsonrpc.sql

-- Load the TAP functions.
begin;

-- Plan the tests.
select plan(1);

insert into "ViewTable1"("Num") values(22);
update "ViewTable1" set "Num"=62 where "ID"=currval('"Table1_ID_seq"')::integer;

-- Run the tests.
select pass('My test passed, w00t!');

-- Finish the tests and clean up.
select * from finish();
rollback;

select setval('"Table1_ID_seq"', 1, false);

\set QUIET 0