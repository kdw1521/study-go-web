package wandoapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	res := httptest.NewRecorder() // response recorder 작성
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// indexHandler(res, req)

	if res.Code != http.StatusOK { // http status check
		t.Fatal("Failed:", res.Code)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if "Hello World" != string(data) { // response body check
		t.Fatal("Response Failed:", string(data))
	}
}

func TestMinjuPathHandler_WithoutName(t *testing.T) {
	res := httptest.NewRecorder() // response recorder 작성
	req := httptest.NewRequest("GET", "/minju", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// minjuHandler(res, req)

	if res.Code != http.StatusOK { // http status check
		t.Fatal("Failed:", res.Code)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if "Minju World" != string(data) { // response body check
		t.Fatal("Response Failed:", string(data))
	}
}

func TestMinjuPathHandler_WithName(t *testing.T) {
	res := httptest.NewRecorder() // response recorder 작성
	req := httptest.NewRequest("GET", "/minju?name=wando", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// minjuHandler(res, req)

	if res.Code != http.StatusOK { // http status check
		t.Fatal("Failed:", res.Code)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if "Minju wando" != string(data) { // response body check
		t.Fatal("Response Failed:", string(data))
	}
}

func TestWandoHandler_WithoutJson(t *testing.T) {
	res := httptest.NewRecorder() // response recorder 작성
	req := httptest.NewRequest("GET", "/wando", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest { // http status check
		t.Fatal("Failed:", res.Code)
	}
}

func TestWandoHandler_WithJson(t *testing.T) {
	res := httptest.NewRecorder() // response recorder 작성
	req := httptest.NewRequest("POST", "/wando", strings.NewReader(`{"first_name":"wando", "last_name":"kim", "email":"kdw@a.a"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusCreated { // http status check
		t.Fatal("Failed:", res.Code)
	}

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	if err != nil {
		t.Fatal("Failed User struct mapping:", err)
	}
	if "wando" != user.FirstName {
		t.Fatal("Not correct firstName:", user.FirstName)
	}
	if "kim" != user.LastName {
		t.Fatal("Not correct lastName:", user.LastName)
	}
	if "kdw@a.a" != user.Email {
		t.Fatal("Not correct email:", user.Email)
	}
}
