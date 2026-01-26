import React from 'react'
import { Routes, Route, Link } from 'react-router-dom'
import Dashboard from './pages/Dashboard'
import Pipelines from './pages/Pipelines'
import PipelineEditor from './pages/PipelineEditor'
import RunLogs from './pages/RunLogs'

export default function App() {
  return (
    <div className="min-h-screen">
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8 flex items-center justify-between">
          <div className="flex items-center gap-4">
            <h1 className="text-lg font-semibold">db-flow</h1>
            <nav className="space-x-4 hidden sm:block">
              <Link to="/" className="text-sm text-slate-600 hover:text-slate-900">Dashboard</Link>
              <Link to="/pipelines" className="text-sm text-slate-600 hover:text-slate-900">Pipelines</Link>
            </nav>
          </div>
          <div>
            <a href="#" className="text-sm text-slate-600 hover:text-slate-900">Account</a>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/pipelines" element={<Pipelines />} />
          <Route path="/pipelines/new" element={<PipelineEditor />} />
          <Route path="/pipelines/:id" element={<PipelineEditor />} />
          <Route path="/runs/:runId" element={<RunLogs />} />
        </Routes>
      </main>
    </div>
  )
}