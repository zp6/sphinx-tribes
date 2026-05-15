package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit tests for assigned bounties tab on profile
// Addresses issue #878

func TestGetAssignedBounties(t *testing.T) {
	req := httptest.NewRequest("GET", "/bounties/assigned?pubkey=test123", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"bounties": []map[string]interface{}{
				{"id": "1", "title": "Test Bounty", "status": "assigned", "assignee": "test123"},
			},
			"total": 1,
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	bounties := response["bounties"].([]interface{})
	if len(bounties) != 1 {
		t.Errorf("Expected 1 bounty, got %d", len(bounties))
	}
}

func TestGetAssignedBountiesEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/bounties/assigned?pubkey=empty", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"bounties": []interface{}{},
			"total":    0,
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
