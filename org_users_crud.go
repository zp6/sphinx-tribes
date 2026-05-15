package sphinxtribes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

// CRUD endpoints for Organizations users
// Addresses issue #622

type OrgUser struct {
	ID         string    `json:"id"`
	OrgUUID    string    `json:"org_uuid"`
	Pubkey     string    `json:"pubkey"`
	Role       string    `json:"role"`
	Created    time.Time `json:"created"`
}

func GetOrgUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	orgUUID := r.URL.Query().Get("org_uuid")
	if orgUUID == "" {
		http.Error(w, "org_uuid required", http.StatusBadRequest)
		return
	}
	rows, err := db.Query("SELECT id, org_uuid, pubkey, role, created FROM org_users WHERE org_uuid = $1", orgUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	users := []OrgUser{}
	for rows.Next() {
		var u OrgUser
		rows.Scan(&u.ID, &u.OrgUUID, &u.Pubkey, &u.Role, &u.Created)
		users = append(users, u)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func AddOrgUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var u OrgUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	u.Created = time.Now()
	_, err := db.Exec("INSERT INTO org_users (org_uuid, pubkey, role, created) VALUES ($1, $2, $3, $4)", u.OrgUUID, u.Pubkey, u.Role, u.Created)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func UpdateOrgUserRole(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var u OrgUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	_, err := db.Exec("UPDATE org_users SET role = $1 WHERE org_uuid = $2 AND pubkey = $3", u.Role, u.OrgUUID, u.Pubkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func DeleteOrgUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	orgUUID := r.URL.Query().Get("org_uuid")
	pubkey := r.URL.Query().Get("pubkey")
	_, err := db.Exec("DELETE FROM org_users WHERE org_uuid = $1 AND pubkey = $2", orgUUID, pubkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
