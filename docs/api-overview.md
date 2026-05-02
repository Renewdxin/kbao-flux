# API Overview

The public API focuses on reviewable contracts.

## Health

- `GET /healthz`

## Agents

- `GET /api/v1/agents`

## Sessions

- `POST /api/v1/sessions`
- `GET /api/v1/sessions/{id}`

## Messages

- `POST /api/v1/messages`
- `GET /api/v1/sessions/{id}/messages`
- `GET /api/v1/sessions/{id}/stream`

## Device bootstrap

- `POST /xiaozhi/ota/`

The public OTA endpoint validates device identity headers and returns mock connection metadata.

The stream endpoint emits deterministic server-sent events to show the public shape of incremental assistant output. It is not the production streaming pipeline.
