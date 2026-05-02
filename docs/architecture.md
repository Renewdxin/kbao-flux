# Architecture Overview

KBAO Flux is organized around a multi-device conversation backend. The public evaluation build keeps only the stable outer contracts.

```text
Device clients
  |
  | HTTP / WebSocket in production
  v
Gateway boundary
  |
  | session, device, and message contracts
  v
Conversation boundary
  |
  | agent and message APIs
  v
Provider boundary
```

## Modules

- Gateway boundary: accepts device bootstrap and session traffic.
- Manager boundary: exposes identity, device, and agent metadata APIs.
- Conversation boundary: owns public session and message records.
- Provider boundary: represented here as mock output only.

## Production-grade concerns represented as boundaries

The private system has to handle latency, interruption, device identity, audio playback state, and agent ownership at the same time. In this public repository those concerns are intentionally represented as API and module boundaries rather than full algorithms.

This lets reviewers evaluate whether the system is shaped for a real hardware product without receiving the implementation that would make the commercial system easy to clone.

## Evaluation build behavior

The mock server uses in-memory data and deterministic responses. It demonstrates request/response shape but does not include production-grade streaming, prompt assembly, provider routing, audio handling, or device coordination.
