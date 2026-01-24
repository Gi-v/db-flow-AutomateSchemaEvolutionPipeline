#!/bin/bash

# 1. Export the current schema from the live DB (ignoring data)
docker exec my-db pg_dump -s -U postgres postgres > live_schema.sql

# 2. Compare it against a "Reference" schema file you keep in Git
if diff live_schema.sql reference_schema.sql > /dev/null; then
  echo "✅ No drift detected. Database matches Git."
else
  echo "❌ DRIFT DETECTED! Someone changed the database manually."
  diff live_schema.sql reference_schema.sql
  exit 1
fi