#!/usr/bin/env bash
# setup_and_run.sh
# Makes the project runnable: fixes filename mismatches & vite config, installs deps,
# starts the backend stub (if present) and the frontend dev server.
#
# Usage:
#   chmod +x setup_and_run.sh
#   ./setup_and_run.sh
set -euo pipefail

ROOT="$(pwd)"
echo "Working directory: $ROOT"

# Helper for printing errors
err() { echo "ERROR: $*" >&2; }

# Ensure we clean up background backend process when script exits or is interrupted
BACKEND_PID_FILE="$ROOT/.backend.pid"
_cleanup() {
  if [ -f "$BACKEND_PID_FILE" ]; then
    PID="$(cat "$BACKEND_PID_FILE" 2>/dev/null || true)"
    if [ -n "$PID" ] && kill -0 "$PID" >/dev/null 2>&1; then
      echo "Stopping backend (pid $PID)..."
      kill "$PID" || true
    fi
    rm -f "$BACKEND_PID_FILE"
  fi
}
trap _cleanup EXIT INT TERM

# 1) Check Node and npm
if ! command -v node >/dev/null 2>&1; then
  err "node is not installed or not on PATH. Install Node.js and re-run this script."
  exit 1
fi
if ! command -v npm >/dev/null 2>&1; then
  err "npm is not installed or not on PATH. Install Node.js (includes npm) and re-run this script."
  exit 1
fi
echo "Node: $(node -v)  npm: $(npm -v)"

# 2) Ensure vite.config.ts content is correct
VITE_CONFIG_FILE="$ROOT/vite.config.ts"
echo "Writing $VITE_CONFIG_FILE ..."
cat > "$VITE_CONFIG_FILE" <<'EOF'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173
  }
})
EOF
echo "Wrote $VITE_CONFIG_FILE"

# 3) Fix filename mismatch: PipelinesEditor -> PipelineEditor
OLD_EDITOR="$ROOT/src/pages/PipelinesEditor.tsx"
NEW_EDITOR="$ROOT/src/pages/PipelineEditor.tsx"
if [ -f "$OLD_EDITOR" ] && [ ! -f "$NEW_EDITOR" ]; then
  echo "Renaming $OLD_EDITOR -> $NEW_EDITOR"
  mkdir -p "$(dirname "$NEW_EDITOR")"
  mv "$OLD_EDITOR" "$NEW_EDITOR"
elif [ -f "$NEW_EDITOR" ]; then
  echo "PipelineEditor exists ($NEW_EDITOR) - ok"
else
  # If neither file exists, create a minimal PipelineEditor so app compiles
  echo "No PipelineEditor found. Creating a minimal $NEW_EDITOR"
  mkdir -p "$(dirname "$NEW_EDITOR")"
  cat > "$NEW_EDITOR" <<'EOF'
import React from 'react'

export default function PipelineEditor() {
  return (
    <div className="max-w-3xl">
      <h2 className="text-2xl font-semibold mb-4">Pipeline Editor (placeholder)</h2>
      <div className="bg-white p-6 rounded shadow">
        <p className="text-slate-600">This is a placeholder editor. Create or open a pipeline to edit it.</p>
      </div>
    </div>
  )
}
EOF
fi

# 4) If backend stub exists, create .env to point frontend to it (do not overwrite)
if [ -d "$ROOT/backend" ] && [ -f "$ROOT/backend/index.js" ]; then
  ENV_FILE="$ROOT/.env"
  if [ ! -f "$ENV_FILE" ]; then
    echo "Creating $ENV_FILE with VITE_API_URL pointing at backend stub"
    cat > "$ENV_FILE" <<'EOF'
VITE_API_URL=http://localhost:4000
EOF
  else
    echo ".env already exists; leaving it untouched"
  fi
fi

# 5) Install root dependencies
cd "$ROOT"
if [ -f package-lock.json ] || [ -f npm-shrinkwrap.json ]; then
  echo "Found lockfile; running npm ci ..."
  npm ci
else
  echo "No lockfile; running npm install ..."
  npm install
fi

# 6) Ensure vite and @vitejs/plugin-react are installed as dev deps
echo "Ensuring vite and @vitejs/plugin-react are installed (devDependencies)..."
npm install -D vite @vitejs/plugin-react --no-audit --no-fund

# 7) Ensure test/dev deps for running tests (idempotent)
echo "Ensuring test/dev dependencies (vitest, @testing-library) are installed..."
npm install -D vitest @testing-library/react @testing-library/jest-dom jsdom --no-audit --no-fund

# 8) Install backend dependencies if a backend stub exists
if [ -d "$ROOT/backend" ] && [ -f "$ROOT/backend/package.json" ]; then
  echo "Installing backend dependencies..."
  (cd "$ROOT/backend" && npm install)
fi

# 9) Start backend (if present)
if [ -d "$ROOT/backend" ] && [ -f "$ROOT/backend/index.js" ]; then
  if [ -f "$BACKEND_PID_FILE" ]; then
    OLD_PID="$(cat "$BACKEND_PID_FILE" 2>/dev/null || true)"
    if [ -n "$OLD_PID" ] && kill -0 "$OLD_PID" >/dev/null 2>&1; then
      echo "Backend already running (pid $OLD_PID)."
    else
      echo "Starting backend..."
      nohup node "$ROOT/backend/index.js" > "$ROOT/.backend.log" 2>&1 &
      BG_PID=$!
      echo "$BG_PID" > "$BACKEND_PID_FILE"
      echo "Backend started (pid $BG_PID). Logs: $ROOT/.backend.log"
    fi
  else
    echo "Starting backend..."
    nohup node "$ROOT/backend/index.js" > "$ROOT/.backend.log" 2>&1 &
    BG_PID=$!
    echo "$BG_PID" > "$BACKEND_PID_FILE"
    echo "Backend started (pid $BG_PID). Logs: $ROOT/.backend.log"
  fi
else
  echo "No backend stub detected; skipping backend start."
fi

# 10) Final checks: run typecheck once (non-fatal)
echo "Running quick typecheck (tsc --noEmit); errors will be shown but do not stop the script."
if command -v npx >/dev/null 2>&1; then
  npx -y tsc --noEmit || true
else
  echo "npx not found; skipping typecheck."
fi

# 11) Start frontend dev server in foreground so user can see logs
echo
echo "Starting frontend dev server: npm run dev"
echo "If you're in GitHub Codespaces or remote container, make sure port 5173 is forwarded."
echo "To stop: Ctrl+C (this will stop the frontend and trigger cleanup to stop backend)."
echo

# Run dev server in foreground (exec so ctrl+c stops it and triggers trap)
exec npm run dev