package http_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t-east/pr-sounds-backend/src/asserts"
	"github.com/t-east/pr-sounds-backend/src/core"
	"github.com/t-east/pr-sounds-backend/src/domains/entities"
	"github.com/t-east/pr-sounds-backend/src/interfaces/controllers"
	"github.com/t-east/pr-sounds-backend/src/interfaces/gateways"
)

// UserName, EmailのあるユーザをPOST -> 201を返すかをテスト
func TestCreateUser(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&entities.User{
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/users", &buf)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Result().Status)
}

func TestGetUser(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}

	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	user, err := ur.Create(u)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%s", user.ID), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
}

//* 存在しないuser.Idを利用
func TestGetUserError(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}

	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	u.ID = core.MakeULID()
	_, err = ur.Create(u)
	if err != nil {
		t.Errorf("Can't create user: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%s", "sdsdfsdfsdfsfd"), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusBadRequest, rec.Code, rec.Result().Status)
}
