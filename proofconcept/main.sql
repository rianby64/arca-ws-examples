
\unset ECHO
\set QUIET 1

create table if not exists "Table1"
(
  "ID" serial,
  "Num" double precision default 0.0,
  constraint "Table1_pkey" primary key ("ID")
);

create or replace view "ViewTable1" as (
  select
    "Table1"."ID" as "ID",
    "Table1"."Num" as "Num"
  from "Table1"
);

create or replace function process_viewtable1()
  returns trigger
  language 'plpgsql' volatile
as $$
declare
  r record;
begin
if tg_op = 'DELETE' then
  for r in (
    select
      'Table1' as "Source",
      lower(tg_op) as "Method",
      row_to_json(t) as "Result",
      true as "Primary"
    from (
      select
        old."ID" as "ID"
    ) t
  ) loop
    perform send_jsonrpc(row_to_json(r));
  end loop;
  return old;
end IF;
if tg_op = 'UPDATE' then
  if (new."Num" <> old."Num") then
    for r in (
      select
        'Table1' as "Source",
        lower(tg_op) as "Method",
        row_to_json(t) as "Result",
        true as "Primary"
      from (
        select
          new."ID" as "ID",
          new."Num" as "Num"
      ) t
    ) loop
      perform send_jsonrpc(row_to_json(r));
    end loop;
  end IF;
  return new;
end IF;
if tg_op = 'INSERT' then
  for r in (
    select
      'Table1' as "Source",
      lower(tg_op) as "Method",
      row_to_json(t) as "Result",
      true as "Primary"
    from (
      select
        new."Num" as "Num"
    ) t
  ) loop
    perform send_jsonrpc(row_to_json(r));
  end loop;
  return new;
end IF;
return null;
end;
$$;

drop trigger if exists "ViewTable1_process" ON "ViewTable1";
create trigger "ViewTable1_process"
instead of insert or update or delete on "ViewTable1"
for each row
execute procedure process_viewtable1();

\set QUIET 0