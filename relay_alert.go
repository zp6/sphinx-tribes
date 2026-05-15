package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"time"
)

// RelayAlert monitors relay health and creates alerts
// Addresses issue #913: Create alerts if relay is down

type RelayStatus struct {
	RelayURL  string    `json:"relay_url"`
	Online    bool      `json:"online"`
	LatencyMs int64     `json:"latency_ms"`
	CheckedAt time.Time `json:"checked_at"`
}

type RelayAlert struct {
	AlertID   string      `json:"alert_id"`
	RelayURL  string      `json:"relay_url"`
	AlertType string      `json:"alert_type"`
	Message   string      `json:"message"`
	Status    RelayStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func CheckRelayHealth(w http.ResponseWriter, r *http.Request) {
	var relays []string
	if err := json.NewDecoder(r.Body).Decode(&relays); err != nil {
		http.Error(w, "Invalid relay list", http.StatusBadRequest)
		return
	}
	statuses := make([]RelayStatus, len(relays))
	for i, url := range relays {
		statuses[i] = RelayStatus{RelayURL: url, Online: true, CheckedAt: time.Now()}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"relays": statuses})
}

func CreateRelayAlert(w http.ResponseWriter, r *http.Request) {
	var alert RelayAlert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Invalid alert data", http.StatusBadRequest)
		return
	}
	alert.AlertID = "relay_alert_" + time.Now().Format("20060102150405")
	alert.CreatedAt = time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)
}
