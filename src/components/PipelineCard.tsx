import React from 'react'

type Pipeline = {
  id: string
  name: string
  db: string
  lastRun: string
  status: string
}

export default function PipelineCard({ pipeline }: { pipeline: Pipeline }) {
  return (
    <div className="bg-white rounded shadow p-4">
      <div className="flex items-start justify-between">
        <div>
          <h3 className="text-lg font-medium">{pipeline.name}</h3>
          <div className="text-sm text-slate-500">
            DB: {pipeline.db} • Last run: {pipeline.lastRun}
          </div>
        </div>
        <div className="text-sm">
          <span className={`px-2 py-1 rounded text-white ${pipeline.status === 'OK' ? 'bg-green-600' : 'bg-red-600'}`}>
            {pipeline.status}
          </span>
        </div>
      </div>

      <div className="mt-4 flex gap-2">
        <button className="px-3 py-1 bg-slate-100 rounded text-sm">Preview</button>
        <button className="px-3 py-1 bg-indigo-600 text-white rounded text-sm">Run</button>
        <button className="px-3 py-1 bg-white border rounded text-sm">History</button>
      </div>
    </div>
  )
}
