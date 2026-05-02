package mockruntime

import (
	"fmt"
	"sync"

	"github.com/Renewdxin/kbao-flux/cmd/mock-server/internal/publicapi"
)

type Store struct {
	mu       sync.Mutex
	nextID   int
	sessions map[string]publicapi.Session
	messages map[string][]publicapi.Message
}

func NewStore() *Store {
	return &Store{
		nextID:   1,
		sessions: make(map[string]publicapi.Session),
		messages: make(map[string][]publicapi.Message),
	}
}

func (s *Store) Agents() []publicapi.Agent {
	return []publicapi.Agent{
		{ID: "guide", Name: "Guide", Description: "Demo assistant for public evaluation."},
		{ID: "companion", Name: "Companion", Description: "Demo companion agent with mock responses."},
	}
}

func (s *Store) CreateSession(deviceID, agentID string) publicapi.Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	session := publicapi.Session{
		ID:       "demo-session",
		DeviceID: deviceID,
		AgentID:  agentID,
	}
	s.sessions[session.ID] = session
	return session
}

func (s *Store) Session(id string) (publicapi.Session, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[id]
	return session, ok
}

func (s *Store) AddMessage(req publicapi.CreateMessageRequest) (publicapi.Message, publicapi.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userMessage := publicapi.Message{
		ID:        s.nextMessageID(),
		SessionID: req.SessionID,
		Role:      "user",
		Text:      req.Text,
	}
	assistantMessage := publicapi.Message{
		ID:        s.nextMessageID(),
		SessionID: req.SessionID,
		Role:      "assistant",
		Text:      "This is a deterministic mock response for event review.",
	}
	s.messages[req.SessionID] = append(s.messages[req.SessionID], userMessage, assistantMessage)
	return userMessage, assistantMessage
}

func (s *Store) Messages(sessionID string) []publicapi.Message {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]publicapi.Message, len(s.messages[sessionID]))
	copy(out, s.messages[sessionID])
	return out
}

func (s *Store) nextMessageID() string {
	id := fmt.Sprintf("msg-%04d", s.nextID)
	s.nextID++
	return id
}
