import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getRun, fetchRunLogs } from '../api/client'

export default function RunLogs() {
  const { runId } = useParams()
  const [logs, setLogs] = useState<string[]>([])
  const [status, setStatus] = useState('pending')

  useEffect(() => {
    if (!runId) return
    let mounted = true
    // initial fetch run metadata
    getRun(runId).then((r) => {
      if (!mounted) return
      setStatus(r.status)
    }).catch(() => {})

    // start polling logs every 1s
    let cursor = 0
    const tick = async () => {
      try {
        const lines = await fetchRunLogs(runId, { after: cursor })
        if (!mounted) return
        if (lines.length > 0) {
          setLogs((l) => [...l, ...lines.map((x: any) => `[${x.level}] ${x.message}`)])
          cursor += lines.length
        }
        // check again for status
        const r = await getRun(runId)
        if (!mounted) return
        setStatus(r.status)
        if (r.status === 'completed' || r.status === 'failed') {
          return // stop polling if finished
        }
      } catch (e) {
        // ignore transient errors
      } finally {
        if (mounted) setTimeout(tick, 1000)
      }
    }
    tick()
    return () => {
      mounted = false
    }
  }, [runId])

  return (
    <div>
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-semibold">Run logs</h2>
        <div className="text-sm text-slate-600">Status: <span className="font-medium">{status}</span></div>
      </div>

      <div className="bg-black text-white rounded p-4 font-mono text-sm h-96 overflow-auto">
        {logs.length === 0 ? <div className="text-slate-400">Waiting for logs…</div> : logs.map((l, i) => <div key={i}>{l}</div>)}
      </div>
    </div>
  )
}