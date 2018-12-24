#!/bin/bash
export PGPASSWORD="arca";
R_FILE="psql -U arca -d arca-1 -w -f"

$R_FILE ./setup-db1.sql

