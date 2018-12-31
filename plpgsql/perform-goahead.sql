
  -- This is the function I want to test
DO
$$
DECLARE
  i bigint=0;
BEGIN
  WHILE i < 2 LOOP
    i = i + 1;
    PERFORM goahead(i);
  END LOOP;
END;
$$;
