# Demo Script

This script gives reviewers a short deterministic path through the public build.

## Start

```bash
go run ./cmd/mock-server
```

## Health check

```bash
curl -s http://127.0.0.1:8088/healthz
```

Expected shape:

```json
{"status":"ok"}
```

## List agents

```bash
curl -s http://127.0.0.1:8088/api/v1/agents
```

## Create a session

```bash
curl -s -X POST http://127.0.0.1:8088/api/v1/sessions \
  -H 'Content-Type: application/json' \
  -d '{"device_id":"demo-device","agent_id":"guide"}'
```

## Send a message

```bash
curl -s -X POST http://127.0.0.1:8088/api/v1/messages \
  -H 'Content-Type: application/json' \
  -d '{"session_id":"demo-session","device_id":"demo-device","text":"hello"}'
```

## Stream mock output

```bash
curl -N 'http://127.0.0.1:8088/api/v1/sessions/demo-session/stream?text=hello'
```

The streaming endpoint emits server-sent events with deterministic mock chunks.

## Device bootstrap

```bash
curl -s -X POST http://127.0.0.1:8088/xiaozhi/ota/ \
  -H 'Device-Id: demo-device' \
  -H 'Client-Id: demo-client'
```
