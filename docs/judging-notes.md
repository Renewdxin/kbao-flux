# Judging Notes

This repository is optimized for event review. It gives judges enough to inspect the architecture, run a demo, and understand the product direction while keeping private commercial implementation out of scope.

## Core idea

KBAO Flux treats hardware devices as first-class participants in a server-side conversation system. The product is not just a chatbot UI. It is a backend for multi-device, persistent, voice-first interaction.

## Differentiators

- Multi-device model: sessions, messages, devices, and agents are server-side resources.
- Hardware-aware API surface: device bootstrap and conversation APIs are separated.
- Reviewable contracts: OpenAPI is included for the public surface.
- Commercial boundary: sensitive algorithms and prompts are excluded rather than obfuscated.
- Runnable demo: the mock server proves request/response and streaming shapes without relying on external vendors.

## Why the repository is intentionally partial

The private system contains production prompt routing, real-time voice pipeline logic, echo handling, dual-device coordination, voiceprint isolation, and provider integration. Those parts are precisely where the commercial value and operational risk live.

For evaluation, this repository exposes the shape of the system and a deterministic mock runtime. That is enough to verify engineering direction without giving away the production implementation.

## Suggested review path

1. Read `README.md` for the project summary and quick start.
2. Read `docs/security-boundary.md` to understand what is intentionally excluded.
3. Inspect `openapi/kbao-flux.public.openapi.yaml`.
4. Run `go test ./...`.
5. Run `go run ./cmd/mock-server` and call the demo endpoints.

## What to look for

- Clear separation between device bootstrap, session ownership, message flow, and provider output.
- Public API design that can support hardware devices instead of only browser clients.
- Explicit handling of disclosure boundaries, which matters for a project intended to become commercial.
