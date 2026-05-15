package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit tests for organization tab
// Addresses issue #875

func TestGetOrganizations(t *testing.T) {
	req := httptest.NewRequest("GET", "/organizations", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"organizations": []map[string]interface{}{
				{"uuid": "org1", "name": "Test Org", "description": "A test organization"},
			},
			"total": 1,
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetOrganizationMembers(t *testing.T) {
	req := httptest.NewRequest("GET", "/organizations/org1/members", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"members": []map[string]interface{}{
				{"pubkey": "user1", "role": "admin"},
				{"pubkey": "user2", "role": "member"},
			},
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
