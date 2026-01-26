import axios from 'axios'

const baseURL = import.meta.env.VITE_API_URL ?? '/api'

const client = axios.create({
  baseURL,
  timeout: 15000
})

// Pipelines
export async function fetchPipelines() {
  const resp = await client.get('/pipelines')
  return resp.data
}

export async function fetchPipeline(id: string) {
  const resp = await client.get(`/pipelines/${id}`)
  return resp.data
}

export async function createPipeline(payload: any) {
  const resp = await client.post(`/pipelines`, payload)
  return resp.data
}

export async function savePipeline(id: string | undefined, payload: any) {
  if (!id) {
    return createPipeline(payload)
  }
  const resp = await client.put(`/pipelines/${id}`, payload)
  return resp.data
}

export async function runPipeline(id: string, body: any) {
  const resp = await client.post(`/pipelines/${id}/run`, body)
  return resp.data
}

// Runs & logs
export async function getRun(runId: string) {
  const resp = await client.get(`/runs/${runId}`)
  return resp.data
}

export async function fetchRunLogs(runId: string, opts?: { after?: number }) {
  const params: any = {}
  if (opts?.after !== undefined) params.after = opts.after
  const resp = await client.get(`/runs/${runId}/logs`, { params })
  return resp.data
}

export default client