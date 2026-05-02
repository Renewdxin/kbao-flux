# KBAO Flux

KBAO Flux 是一个面向赛事评审的多设备对话后端公开版本。它展示系统的公开边界：设备接入、Agent 选择、会话与消息 API、OTA 风格的设备启动配置，以及一个可运行的 mock 实时响应路径。

本仓库是评审版，不是完整商业实现。生产级实时对话编排、prompt、回声处理、双硬件协同、声纹实现、真实 provider adapter、部署自动化和商业运营逻辑均未公开。

## 项目价值

大多数语音助手 demo 默认是“一个用户、一个设备、一个会话”。KBAO Flux 面向的是更接近真实硬件产品的问题：一个用户可以拥有多个设备，设备可以共享服务端会话、Agent、消息历史和设备状态。

这个公开版本用于让评委确认系统方向、工程结构和 API 设计，同时保护未来商业化所需的核心实现。

## 本仓库包含

- 高层架构说明和模块边界。
- 设备启动、会话、消息、Agent、健康检查等公开 API 契约。
- 不含真实密钥和生产地址的示例配置。
- 一个可运行的 mock server，不依赖外部 LLM/TTS/ASR 服务。
- 一个 mock SSE 流式输出端点，用于展示公开流式接口形态。

## 本仓库不包含

- system prompt、agent prompt、prompt 路由和上下文组装逻辑。
- 低延迟 LLM 到 TTS 的流式管线、播放队列、打断恢复和缓冲策略。
- 跨设备回声处理、VAD 事件关联和双硬件协同。
- 声纹注册、识别、验证、阈值、向量库和模型隔离实现。
- 真实 provider adapter、生产部署脚本、密钥和商业后台逻辑。

## 快速运行

```bash
go run ./cmd/mock-server
```

另开一个终端：

```bash
curl -s http://127.0.0.1:8088/healthz
curl -s http://127.0.0.1:8088/api/v1/agents
curl -s -X POST http://127.0.0.1:8088/api/v1/sessions \
  -H 'Content-Type: application/json' \
  -d '{"device_id":"demo-device","agent_id":"guide"}'
curl -N 'http://127.0.0.1:8088/api/v1/sessions/demo-session/stream?text=hello'
```

## 评审建议

建议按以下顺序查看：

1. `docs/judging-notes.md`
2. `docs/security-boundary.md`
3. `openapi/kbao-flux.public.openapi.yaml`
4. `docs/demo-script.md`

## 许可证

本仓库仅用于赛事评审。详见 `LICENSE`。
