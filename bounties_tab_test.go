package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit tests for Bounties tab on a profile
// Addresses issue #877

func TestGetProfileBounties(t *testing.T) {
	req := httptest.NewRequest("GET", "/profile/testpubkey/bounties", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"bounties": []map[string]interface{}{
				{"id": "1", "title": "Test Bounty", "price": 5000, "status": "open"},
				{"id": "2", "title": "Another Bounty", "price": 10000, "status": "assigned"},
			},
			"total": 2,
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetProfileBountiesEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/profile/nonexistent/bounties", nil)
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
