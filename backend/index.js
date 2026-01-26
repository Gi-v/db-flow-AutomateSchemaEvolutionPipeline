const express = require('express')
const cors = require('cors')
const { v4: uuidv4 } = require('uuid')

const app = express()
app.use(cors())
app.use(express.json())

/**
 * In-memory stores for pipelines & runs. For a real backend, replace with DB.
 */
const pipelines = {}
const runs = {}

/**
 * Seed a sample pipeline
 */
const sampleId = 'p-1'
pipelines[sampleId] = {
  id: sampleId,
  name: 'Create Products Table',
  db: 'mysql',
  repo: '',
  steps: []
}

// List pipelines
app.get('/api/pipelines', (req, res) => {
  res.json(Object.values(pipelines))
})

// Get pipeline
app.get('/api/pipelines/:id', (req, res) => {
  const p = pipelines[req.params.id]
  if (!p) return res.status(404).json({ message: 'Not found' })
  res.json(p)
})

// Create pipeline
app.post('/api/pipelines', (req, res) => {
  const id = uuidv4()
  const payload = { id, ...req.body }
  pipelines[id] = payload
  res.status(201).json(payload)
})

// Update pipeline
app.put('/api/pipelines/:id', (req, res) => {
  const id = req.params.id
  if (!pipelines[id]) return res.status(404).json({ message: 'Not found' })
  pipelines[id] = { ...pipelines[id], ...req.body }
  res.json(pipelines[id])
})

// Start a run (simulate a run; returns run id)
app.post('/api/pipelines/:id/run', (req, res) => {
  const runId = uuidv4()
  const pipelineId = req.params.id
  const run = {
    id: runId,
    pipelineId,
    status: 'running',
    startedAt: new Date().toISOString(),
    logs: []
  }
  runs[runId] = run

  // Simulate logs over time
  const messages = [
    { level: 'INFO', message: 'Starting pipeline' },
    { level: 'INFO', message: 'Extracting schema' },
    { level: 'INFO', message: 'Generating migration SQL' },
    { level: 'WARN', message: 'Non-nullable column added without default' },
    { level: 'INFO', message: 'Applying migration' },
    { level: 'INFO', message: 'Migration complete' }
  ]

  let i = 0
  const t = setInterval(() => {
    if (!runs[runId]) {
      clearInterval(t)
      return
    }
    if (i < messages.length) {
      runs[runId].logs.push(messages[i])
      i++
    } else {
      runs[runId].status = 'completed'
      runs[runId].finishedAt = new Date().toISOString()
      clearInterval(t)
    }
  }, 700)

  res.status(201).json({ id: runId })
})

// Get run metadata
app.get('/api/runs/:runId', (req, res) => {
  const r = runs[req.params.runId]
  if (!r) return res.status(404).json({ message: 'Not found' })
  res.json({ id: r.id, status: r.status, startedAt: r.startedAt, finishedAt: r.finishedAt })
})

// Get run logs (supports ?after=N as cursor)
app.get('/api/runs/:runId/logs', (req, res) => {
  const r = runs[req.params.runId]
  if (!r) return res.status(404).json({ message: 'Not found' })
  const after = parseInt(req.query.after || '0', 10)
  const slice = r.logs.slice(after)
  res.json(slice)
})

const port = process.env.PORT || 4000
app.listen(port, () => {
  console.log(`Backend stub listening on http://localhost:${port}`)
})