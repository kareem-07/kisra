package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

type Order struct {
	Name string `json:"name"`
	Quantity string `json:"quantity"`
	Date string `json:"date"`
	PaymentMethod string `json:"payment_method"`
	Location string `json:"location"`
}


func POST_order(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<15)
	var payload Order
	decoder := json.NewDecoder(r.Body)
	if e := decoder.Decode(&payload); e!=nil{
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	if payload.Name == "" {
		http.Error(w, "Mandatory form value 'name' not provided", http.StatusBadRequest)
		return
	}
	if payload.Quantity == "" {
		http.Error(w, "Mandatory form value 'quantity' not provided", http.StatusBadRequest)
		return
	}
	if payload.Date == "" {
		http.Error(w, "Mandatory form value 'date' not provided", http.StatusBadRequest)
		return
	}
	if payload.PaymentMethod == "" {
		http.Error(w, "Mandatory form value 'payment_method' not provided", http.StatusBadRequest)
		return
	}
	if payload.Location == "" {
		http.Error(w, "Mandatory form value 'location' not provided", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All clear"))

}

func main() {
	rootDir := ".."
	frontendDir := filepath.Join(rootDir, "frontend")
	fs := http.FileServer(http.Dir(frontendDir))

	http.Handle("/", fs)
	http.HandleFunc("POST /order", func(w http.ResponseWriter, r *http.Request) {POST_order(w, r)})
	
	log.Println("Starting server")
	e := http.ListenAndServe(":8080", nil); if e!=nil{log.Fatal(e)}
}




// import (
// 	"database/sql"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// 	"path/filepath"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func POST_family(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	bodyBytes, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read body", http.StatusInternalServerError)
// 		return
// 	}

// 	var payload map[string]any
// 	if err = json.Unmarshal(bodyBytes, &payload); err != nil {
// 		http.Error(w, "Invalid json given", http.StatusBadRequest)
// 		return
// 	}

// 	username, _ := payload["username"].(string)
// 	familyData, _ := json.Marshal(payload["family"])
// 	stmt := `INSERT INTO users (username, family) VALUES (?, ?)
// 					ON DUPLICATE KEY UPDATE family = VALUES(family)`
// 	_, err = db.Exec(stmt, string(username), string(familyData))
// 	if err != nil {
// 		http.Error(w, "DB insert failed: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Data written successfully"))
// 	log.Println("Inserted data to the user", username)
// }

// func GET_family(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, "Failed to parse form", http.StatusBadRequest)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	if username == "" {
// 		http.Error(w, "Username not provided", http.StatusBadRequest)
// 		return
// 	}

// 	stmt := `SELECT family FROM users WHERE username = ?`
// 	var familyData string
// 	if err := db.QueryRow(stmt, username).Scan(&familyData); err != nil {
// 		http.Error(w, "DB select failed: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte(familyData))

// 	log.Println("Got the username", username)
// }

// func main() {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/family_tree")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	rootDir := "../.."
// 	distDir := filepath.Join(rootDir, "frontend", "dist")
// 	fs := http.FileServer(http.Dir(distDir))

// 	http.Handle("/", fs)
// 	http.HandleFunc("GET /family", func(w http.ResponseWriter, r *http.Request) {GET_family(w, r, db)})
// 	http.HandleFunc("POST /family", func(w http.ResponseWriter, r *http.Request) {POST_family(w, r, db)})

// 	log.Println("Server is running from http://localhost:8080")
	
// 	err = http.ListenAndServe(":8080", nil)	
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }