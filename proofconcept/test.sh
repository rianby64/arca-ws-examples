#!/bin/bash
export PGPASSWORD="heka";
R_FILE="psql -U heka -d heka-0 -w -f"

$R_FILE ./test.sql
