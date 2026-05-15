package sphinxtribes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TicketOwnerMigration handles migrating ticket ownership from V1 to V2 accounts
// Addresses issue #1919: Update Ticket Owner from V1 Account to V2

type MigrationRequest struct {
	TicketID string `json:"ticket_id"`
	V1Owner  string `json:"v1_owner"`
	V2Owner  string `json:"v2_owner"`
}

type MigrationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// MigrateTicketOwner updates the ticket owner from V1 to V2 account
func MigrateTicketOwner(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req MigrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.TicketID == "" || req.V1Owner == "" || req.V2Owner == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	result, err := db.Exec(
		"UPDATE tickets SET owner_pubkey = $1, updated_at = $2 WHERE id = $3 AND owner_pubkey = $4",
		req.V2Owner, time.Now(), req.TicketID, req.V1Owner,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Migration failed: %v", err), http.StatusInternalServerError)
		return
	}
	rows, _ := result.RowsAffected()
	response := MigrationResponse{
		Success: rows > 0,
		Message: fmt.Sprintf("Migrated %d ticket(s) from V1 to V2", rows),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
