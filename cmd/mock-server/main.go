package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Renewdxin/kbao-flux/cmd/mock-server/internal/mockruntime"
	"github.com/Renewdxin/kbao-flux/cmd/mock-server/internal/publicapi"
)

func main() {
	store := mockruntime.NewStore()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /api/v1/agents", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"data": store.Agents()})
	})

	mux.HandleFunc("POST /api/v1/sessions", func(w http.ResponseWriter, r *http.Request) {
		var req publicapi.CreateSessionRequest
		if !decodeJSON(w, r, &req) {
			return
		}
		if req.DeviceID == "" || req.AgentID == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "device_id and agent_id are required"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"data": store.CreateSession(req.DeviceID, req.AgentID)})
	})

	mux.HandleFunc("GET /api/v1/sessions/{id}", func(w http.ResponseWriter, r *http.Request) {
		session, ok := store.Session(r.PathValue("id"))
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "session not found"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"data": session})
	})

	mux.HandleFunc("POST /api/v1/messages", func(w http.ResponseWriter, r *http.Request) {
		var req publicapi.CreateMessageRequest
		if !decodeJSON(w, r, &req) {
			return
		}
		if req.SessionID == "" || req.DeviceID == "" || req.Text == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "session_id, device_id, and text are required"})
			return
		}
		userMessage, assistantMessage := store.AddMessage(req)
		writeJSON(w, http.StatusOK, map[string]any{
			"data": map[string]any{
				"user_message":      userMessage,
				"assistant_message": assistantMessage,
			},
		})
	})

	mux.HandleFunc("GET /api/v1/sessions/{id}/messages", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"data": store.Messages(r.PathValue("id"))})
	})

	mux.HandleFunc("GET /api/v1/sessions/{id}/stream", func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.PathValue("id")
		if _, ok := store.Session(sessionID); !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "session not found"})
			return
		}
		writeSSE(w, []string{
			"This ",
			"is ",
			"a ",
			"mock ",
			"stream ",
			"for ",
			"event ",
			"review.",
		})
	})

	mux.HandleFunc("POST /xiaozhi/ota/", func(w http.ResponseWriter, r *http.Request) {
		deviceID := strings.TrimSpace(r.Header.Get("Device-Id"))
		clientID := strings.TrimSpace(r.Header.Get("Client-Id"))
		if deviceID == "" || clientID == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Device-Id and Client-Id are required"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"mode": "mock",
			"websocket": map[string]string{
				"url": "ws://127.0.0.1:8088/mock-ws",
			},
		})
	})

	addr := os.Getenv("KBAO_FLUX_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8088"
	}

	log.Printf("mock server listening on http://%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func writeSSE(w http.ResponseWriter, chunks []string) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	flusher, ok := w.(http.Flusher)
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "streaming unsupported"})
		return
	}

	for _, chunk := range chunks {
		if _, err := w.Write([]byte("event: delta\n")); err != nil {
			return
		}
		if _, err := w.Write([]byte("data: " + chunk + "\n\n")); err != nil {
			return
		}
		flusher.Flush()
		time.Sleep(80 * time.Millisecond)
	}
	_, _ = w.Write([]byte("event: done\ndata: {}\n\n"))
	flusher.Flush()
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("write response: %v", err)
	}
}
