#!/bin/bash
export PGPASSWORD="arca";
R_FILE="psql -U arca -d arca-4 -w -f"

$R_FILE ./load-test-db4.sql