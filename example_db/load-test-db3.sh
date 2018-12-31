#!/bin/bash
export PGPASSWORD="arca";
R_FILE="psql -U arca -d arca-3 -w -f"

$R_FILE ./perform-goahead.sql