import React from 'react'
import PipelineCard from '../components/PipelineCard'

const mockPipelines = [
  { id: 'p-1', name: 'Create Products Table', db: 'mysql', lastRun: '2026-01-25', status: 'OK' },
  { id: 'p-2', name: 'Add price index', db: 'mysql', lastRun: '2026-01-20', status: 'Failed' }
]

export default function Dashboard() {
  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-2xl font-semibold">Pipelines</h2>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Create pipeline</button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {mockPipelines.map(p => (
          <PipelineCard key={p.id} pipeline={p} />
        ))}
      </div>
    </div>
  )
}
