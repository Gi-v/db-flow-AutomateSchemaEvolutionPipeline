import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import PipelineCard from '../components/PipelineCard'
import { fetchPipelines, createPipeline } from '../api/client'

export default function Pipelines() {
  const [pipelines, setPipelines] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [creating, setCreating] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    let mounted = true
    fetchPipelines()
      .then((data) => {
        if (!mounted) return
        setPipelines(data)
      })
      .catch(() => setPipelines([]))
      .finally(() => setLoading(false))
    return () => {
      mounted = false
    }
  }, [])

  async function onCreateQuick() {
    setCreating(true)
    try {
      const newPipeline = await createPipeline({
        name: `New pipeline ${Date.now()}`,
        db: 'mysql',
        repo: '',
        steps: []
      })
      navigate(`/pipelines/${newPipeline.id}`)
    } catch (e) {
      alert('Failed to create pipeline')
    } finally {
      setCreating(false)
    }
  }

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-2xl font-semibold">Pipelines</h2>
        <div className="flex items-center gap-2">
          <Link to="/pipelines/new" className="bg-white border px-4 py-2 rounded hover:bg-slate-50 text-sm">Create pipeline</Link>
          <button
            className="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700 text-sm"
            onClick={onCreateQuick}
            disabled={creating}
          >
            {creating ? 'Creating…' : 'Quick create'}
          </button>
        </div>
      </div>

      {loading ? (
        <div className="text-slate-500">Loading pipelines…</div>
      ) : pipelines.length === 0 ? (
        <div className="text-slate-600">No pipelines yet. Create one to get started.</div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {pipelines.map((p) => (
            <PipelineCard key={p.id} pipeline={p} />
          ))}
        </div>
      )}
    </div>
  )
}