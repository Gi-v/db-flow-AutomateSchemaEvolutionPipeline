import React, { useEffect, useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { fetchPipeline, savePipeline, runPipeline } from '../api/client'

export default function PipelineEditor() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [name, setName] = useState('')
  const [db, setDb] = useState('mysql')
  const [repo, setRepo] = useState('')
  const [saving, setSaving] = useState(false)
  const [running, setRunning] = useState(false)

  useEffect(() => {
    if (!id) return
    fetchPipeline(id)
      .then((p) => {
        setName(p.name)
        setDb(p.db)
        setRepo(p.repo ?? '')
      })
      .catch(() => {})
  }, [id])

  async function onSave(e: React.FormEvent) {
    e.preventDefault()
    setSaving(true)
    try {
      const saved = await savePipeline(id, { name, db, repo })
      navigate(`/pipelines/${saved.id}`)
    } catch (err) {
      alert('Failed to save pipeline')
    } finally {
      setSaving(false)
    }
  }

  async function onRun() {
    setRunning(true)
    try {
      const run = await runPipeline(id ?? 'new', { env: 'dev', mode: 'dry-run' })
      navigate(`/runs/${run.id}`)
    } catch (err) {
      alert('Failed to start run')
    } finally {
      setRunning(false)
    }
  }

  return (
    <div className="max-w-3xl">
      <h2 className="text-2xl font-semibold mb-4">{id ? 'Edit Pipeline' : 'New Pipeline'}</h2>
      <form onSubmit={onSave} className="space-y-4 bg-white p-6 rounded shadow">
        <div>
          <label className="block text-sm font-medium text-slate-700">Name</label>
          <input required value={name} onChange={(e) => setName(e.target.value)} className="mt-1 block w-full border rounded px-3 py-2" />
        </div>

        <div>
          <label className="block text-sm font-medium text-slate-700">Database</label>
          <select value={db} onChange={(e) => setDb(e.target.value)} className="mt-1 block w-48 border rounded px-3 py-2">
            <option value="mysql">MySQL</option>
            <option value="postgres">Postgres</option>
          </select>
        </div>

        <div>
          <label className="block text-sm font-medium text-slate-700">Repository (optional)</label>
          <input value={repo} onChange={(e) => setRepo(e.target.value)} className="mt-1 block w-full border rounded px-3 py-2" />
        </div>

        <div className="flex gap-2 justify-end">
          <button type="button" onClick={() => window.history.back()} className="px-4 py-2 bg-white border rounded">Cancel</button>
          <button type="submit" className="px-4 py-2 bg-indigo-600 text-white rounded" disabled={saving}>{saving ? 'Saving…' : 'Save'}</button>
          <button type="button" onClick={onRun} className="px-4 py-2 bg-green-600 text-white rounded" disabled={running}>{running ? 'Starting…' : 'Run'}</button>
        </div>
      </form>
    </div>
  )
}