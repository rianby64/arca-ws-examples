
\unset ECHO
\set QUIET 1

create or replace function send_jsonrpc(request json)
  returns void
  language 'plpgsql' volatile
as $$
declare
  source text;
  result json;
  method text;
  r record;
  query text;
  ID text;
begin
  source := request ->> 'Source';
  method := request ->> 'Method';
  result := request ->> 'Result';

  if source = 'Table1' then
    if Method = 'insert' then
      insert into "Table1"("Num")
        values ((result ->> 'Num')::integer);
    end if;
    if Method = 'update' then
      query := 'UPDATE "Table1" SET ';
      for r in
        select "key", "value" from json_each(result)
      loop
        if r.key = 'ID' then
          ID := r.value;
          continue;
        end if;
        query = query || format('"%s"=%s, ', r.key, r.value);
      end loop;
      query = format('%s WHERE "ID"=%s', left(query, length(query) - 2), ID);
      raise notice '%', query;
      execute query;
    end if;
  end if;
end;
$$;

\set QUIET 0