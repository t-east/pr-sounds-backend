package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/t-east/pr-sounds-backend/src/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/t-east/pr-sounds-backend/src/domains/entities"
	"github.com/t-east/pr-sounds-backend/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func allowOptionsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	db, err := LoadTestDB()
	if err != nil {
		log.Fatalf("Can't get DB. %+v", err)
	}
	err = allowOptionsMiddleware(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var head string
	_, r.URL.Path = core.ShiftPath(r.URL.Path)
	head, r.URL.Path = core.ShiftPath(r.URL.Path)
	switch head {
	case "users":
		uc := controllers.LoadUserController(db)
		uc.Dispatch(w, r)
	default:
		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
	}
}

// Serve はserverを起動させます．
func Serve() {
	sm := http.NewServeMux()
	sm.Handle("/api/", http.HandlerFunc(ServerHandlerPublic))
	err := http.ListenAndServe(":4000", sm)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
