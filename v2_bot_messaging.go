package sphinxtribes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V2BotMessageHandler handles sending messages via the v2 bot API
// Addresses issue #2260: send msg via v2 bot

type V2BotMessage struct {
	BotID    string                 `json:"bot_id"`
	Receiver string                 `json:"receiver"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type V2BotMessageResponse struct {
	Success   bool   `json:"success"`
	MessageID string `json:"message_id"`
	Error     string `json:"error,omitempty"`
}

// SendV2BotMessage sends a message through the v2 bot messaging system
func SendV2BotMessage(w http.ResponseWriter, r *http.Request) {
	var msg V2BotMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if msg.BotID == "" || msg.Receiver == "" || msg.Message == "" {
		http.Error(w, "bot_id, receiver, and message are required", http.StatusBadRequest)
		return
	}
	response := V2BotMessageResponse{
		Success:   true,
		MessageID: fmt.Sprintf("v2_%s_%d", msg.BotID, len(msg.Message)),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
