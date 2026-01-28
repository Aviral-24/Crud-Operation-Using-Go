package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud/db"
	"crud/models"
)
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		result, err := db.DB.Exec(
			"INSERT INTO users (name, email) VALUES (?, ?)",
			user.Name, user.Email,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		id, _ := result.LastInsertId()
		user.ID = int(id)
		json.NewEncoder(w).Encode(user)

	case http.MethodGet:
		rows, err := db.DB.Query("SELECT id, name, email FROM users")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var u models.User
			rows.Scan(&u.ID, &u.Name, &u.Email)
			users = append(users, u)
		}

		json.NewEncoder(w).Encode(users)

	default:
		http.Error(w, "Method not allowed", 405)
	}
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, _ := strconv.Atoi(idStr)

	switch r.Method {

	case http.MethodPut:
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		_, err := db.DB.Exec(
			"UPDATE users SET name=?, email=? WHERE id=?",
			user.Name, user.Email, id,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		user.ID = id
		json.NewEncoder(w).Encode(user)

	case http.MethodDelete:
		_, err := db.DB.Exec("DELETE FROM users WHERE id=?", id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("User deleted"))

	default:
		http.Error(w, "Method not allowed", 405)
	}
}
