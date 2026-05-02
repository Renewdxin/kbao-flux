# Public Disclosure Boundary

This repository is a source-available evaluation build. It is designed to show engineering structure and public API shape without disclosing proprietary implementation details.

## Excluded by design

- Production prompts and prompt assembly.
- Conversation orchestration internals, including turn-taking, interruption recovery, context compaction, and tool-routing behavior.
- Real-time audio pipeline details, including sentence splitting, streaming TTS, playback scheduling, and latency control.
- Echo suppression and cross-device VAD correlation.
- Dual-hardware conversation coordination and speaker attribution rules.
- Voiceprint enrollment, verification, identification, vector storage, thresholds, and model isolation implementation.
- Real LLM/TTS/ASR provider adapters.
- Production deployment scripts, environment variable maps, credentials, domains, and infrastructure notes.

## Included for review

- High-level module boundaries.
- Public API contracts.
- A deterministic mock server.
- Example configuration with no real secrets.

The mock server is intentionally not a drop-in replacement for the private production system.
