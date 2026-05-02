package publicapi

type Agent struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSessionRequest struct {
	DeviceID string `json:"device_id"`
	AgentID  string `json:"agent_id"`
}

type Session struct {
	ID       string `json:"id"`
	DeviceID string `json:"device_id"`
	AgentID  string `json:"agent_id"`
}

type CreateMessageRequest struct {
	SessionID string `json:"session_id"`
	DeviceID  string `json:"device_id"`
	Text      string `json:"text"`
}

type Message struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Role      string `json:"role"`
	Text      string `json:"text"`
}
