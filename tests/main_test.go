package tests

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/viggin543/go_http_server/bootstrap"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	err := os.Setenv("DB_HOST", "127.0.0.1")
	if err != nil {
		panic("failed to init tests")
	}
	ts = httptest.NewServer(bootstrap.SetupServer())
	defer ts.Close()
	os.Exit(m.Run())
}

func TestPingRoute(t *testing.T) {

	resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	assertRespHeaders(t, err, resp)
	contentLen := contentLength(resp)
	bytes := parseResponse(t, contentLen, resp)
	type pong struct {
		Message string `json:"message"`
	}
	var da pong
	_ = json.Unmarshal(bytes, &da)
	if da.Message != "pong" {
		t.Fatalf("expected pong")
	}

}

func contentLength(resp *http.Response) int {
	contentLen, _ := strconv.Atoi(resp.Header.Get("content-length"))
	return contentLen
}

func parseResponse(t *testing.T, contentLen int, resp *http.Response) []byte {
	var bytes = make([]byte, contentLen)
	n, _ := resp.Body.Read(bytes)
	if n == 0 {
		t.Fatalf("json resp expected")
	}
	return bytes
}


func Test_fruits(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/fruits", ts.URL))
	assertRespHeaders(t, err, resp)
	contentLen := contentLength(resp)
	bytes := parseResponse(t, contentLen, resp)

	var frutas []bootstrap.Fruit
	err = json.Unmarshal(bytes, &frutas)
	if err != nil {
		t.Log(frutas)
		t.Fatalf("expected a frutas resp")
	}
}


func assertRespHeaders(t *testing.T, err error, resp *http.Response) {
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	val, ok := resp.Header["Content-Type"]
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

