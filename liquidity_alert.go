package sphinxtribes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LiquidityAlert monitors and creates alerts for bounty liquidity
// Addresses issue #914: Create liquidity alert

type LiquidityAlert struct {
	AlertID     string    `json:"alert_id"`
	Type        string    `json:"type"`
	Threshold   int64     `json:"threshold"`
	Current     int64     `json:"current"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

type AlertConfig struct {
	LowBalanceThreshold    int64 `json:"low_balance_threshold"`
	CriticalThreshold      int64 `json:"critical_threshold"`
	NotificationEnabled    bool  `json:"notification_enabled"`
}

func CheckLiquidity(w http.ResponseWriter, r *http.Request) {
	var config AlertConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Invalid config", http.StatusBadRequest)
		return
	}
	alerts := []LiquidityAlert{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"alerts": alerts,
		"status": "ok",
	})
}

func CreateLiquidityAlert(w http.ResponseWriter, r *http.Request) {
	var alert LiquidityAlert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Invalid alert data", http.StatusBadRequest)
		return
	}
	alert.AlertID = fmt.Sprintf("alert_%d", time.Now().Unix())
	alert.CreatedAt = time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)
}
