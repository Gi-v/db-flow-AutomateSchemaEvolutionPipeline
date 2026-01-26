# db-flow-AutomateSchemaEvolutionPipeline — React + Tailwind Starter

This repository was extended with a minimal React + TypeScript + Vite + Tailwind starter UI tailored for the db-flow Automate Schema Evolution project. It's a scaffold you can iterate on: a small app shell, example pipeline card component, API client placeholder, and a Docker Compose file to run a local MySQL instance for development.

Quick overview
- Frontend: React + TypeScript + Vite + Tailwind CSS
- API client: Axios + VITE_API_URL env var
- Local DB (optional for dev): MySQL via docker-compose.yml
- Example SQL migrations for MySQL are under sql/mysql/

Getting started (frontend)
1. Install dependencies
   - npm install

2. Run development server
   - npm run dev
   - Open http://localhost:5173

3. Build for production
   - npm run build
   - npm run preview

Environment variables
- VITE_API_URL — base URL for backend API (defaults to `/api`)

Start local MySQL (optional)
- docker compose up -d
- MySQL: root / password, database: dbflow
- Flyway or your migration runner can point to: jdbc:mysql://localhost:3306/dbflow

Where to extend
- src/pages/ — add new pages (Pipelines list, Pipeline editor, Runs)
- src/components/ — reusable UI components
- src/api/client.ts — central place to call backend endpoints
- sql/mysql/ — example migrations (used by Flyway or other runner for tests)

Notes about repository
- This starter intentionally focuses on frontend tooling and local DB for development. Integrate with your existing db-flow backend by implementing API endpoints and setting VITE_API_URL.
