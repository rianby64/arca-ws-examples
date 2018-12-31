#!/bin/bash
export PGPASSWORD="arca";
R_FILE="psql -U arca -d arca-2 -w -f"

$R_FILE ./setup-db2.sql