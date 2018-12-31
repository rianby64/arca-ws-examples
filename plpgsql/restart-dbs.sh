#!/bin/bash
export PGPASSWORD="postgres";
R_FILE="psql -U postgres -d postgres -w -f"

$R_FILE ./restart-dbs.sql