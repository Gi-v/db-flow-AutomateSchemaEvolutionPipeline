import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import { expect, test } from '@jest/globals'
import { MemoryRouter } from 'react-router-dom'
import App from '../App'

test('renders dashboard title', () => {
  render(
    <MemoryRouter initialEntries={['/']}>
      <App />
    </MemoryRouter>
  )
  expect(screen.getByText(/Pipelines/i)).toBeInTheDocument()
})