# KBAO Flux

[中文说明](./README_zh.md)

KBAO Flux is a multi-device conversational backend demo for event review. It shows the public shape of the system: device sessions, agent selection, conversation APIs, OTA-style device bootstrap, and a mock real-time response path.

This repository is intentionally scoped as an evaluation build. It does not include proprietary real-time dialogue orchestration, production prompts, echo cancellation, dual-hardware coordination, voiceprint implementation, provider adapters, deployment automation, or commercial operating logic.

## Submission Info

- Event tag: `#Flux南客松S2`
- Repository: `https://github.com/Renewdxin/kbao-flux`
- Demo URL / package: `<OSS_DEMO_URL>`
- Commercial disclosure: this is a partial source-available evaluation repository. Core commercial implementation remains private.

## Why it matters

Most voice assistant demos assume one account, one device, and one active conversation. KBAO Flux is designed around the harder case: one user can own multiple hardware devices, move between them, and keep server-side sessions, agents, and message history consistent.

This public build demonstrates the contract surface and review path without exposing the private implementation that makes the production system commercially defensible.

## Tech Stack

| Area | Selection | Why |
| --- | --- | --- |
| Language | Go 1.22 | Small deployment footprint, strong concurrency model, simple static binaries. |
| HTTP framework | Go standard `net/http` | Keeps the public evaluation server dependency-light and easy to audit. |
| API contract | OpenAPI 3.1 | Machine-readable review contract for device, session, message, and bootstrap APIs. |
| Streaming demo | Server-Sent Events | Demonstrates incremental assistant output shape without exposing production streaming internals. |
| Runtime state | In-memory mock store | Deterministic review flow with no database or external service dependency. |
| Config | YAML example | Documents intended runtime knobs without exposing production secrets. |
| CI | GitHub Actions | Runs Go checks on repository changes. |
| Validation | `go test`, `go vet`, Postman OpenAPI lint | Covers compile-level correctness and API spec validity. |

## What is included

- Public architecture notes and module boundaries.
- Public API contracts for device bootstrap, sessions, messages, agents, and health checks.
- Example configuration with fake local values only.
- A runnable mock server that returns deterministic responses without external LLM/TTS/ASR providers.
- A mock server-sent events endpoint that demonstrates streaming API shape without exposing real latency-control logic.
- Interface-level contracts that describe integration points without exposing production internals.

## What is not included

- System prompts, agent prompts, prompt routing, or context assembly logic.
- Low-latency LLM-to-TTS streaming pipeline, playback queue, buffering, or interruption handling.
- Cross-device echo suppression, VAD event correlation, or hardware-to-hardware conversation logic.
- Voiceprint model isolation, thresholds, embedding storage, or vector search implementation.
- Real provider adapters, production deployment scripts, secret handling, or commercial admin workflows.

## Quick start

```bash
go run ./cmd/mock-server
```

or:

```bash
make run
```

Then try:

```bash
curl -s http://127.0.0.1:8088/healthz
curl -s http://127.0.0.1:8088/api/v1/agents
curl -s -X POST http://127.0.0.1:8088/api/v1/sessions \
  -H 'Content-Type: application/json' \
  -d '{"device_id":"demo-device","agent_id":"guide"}'
curl -s -X POST http://127.0.0.1:8088/api/v1/messages \
  -H 'Content-Type: application/json' \
  -d '{"session_id":"demo-session","device_id":"demo-device","text":"hello"}'
curl -N 'http://127.0.0.1:8088/api/v1/sessions/demo-session/stream?text=hello'
```

## Repository layout

```text
cmd/mock-server/          Runnable mock HTTP server
configs/                  Example-only configuration
docs/                     Public architecture and disclosure boundary
openapi/                  Public API contracts
```

For reviewers, start with `docs/judging-notes.md`, then run the quick-start commands above.

## Checks

```bash
make test
make vet
```

## License

This repository is provided for event evaluation only. See `LICENSE`.
