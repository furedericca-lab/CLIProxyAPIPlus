package kimi

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v7/internal/config"
)

func TestNewDeviceFlowClientWithDeviceIDAndProxyURL_OverrideDirectDisablesProxy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := &config.Config{SDKConfig: config.SDKConfig{ProxyURL: "http://127.0.0.1:1"}}
	client := NewDeviceFlowClientWithDeviceIDAndProxyURL(cfg, "device-1", "direct")

	resp, errDo := client.httpClient.Get(server.URL)
	if errDo != nil {
		t.Fatalf("client.Get returned error: %v", errDo)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}

func TestNewDeviceFlowClientWithDeviceIDAndProxyURL_OverrideProxyTakesPrecedence(t *testing.T) {
	var proxyCalls int32
	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&proxyCalls, 1)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer proxyServer.Close()

	cfg := &config.Config{SDKConfig: config.SDKConfig{ProxyURL: "http://global.example.com:8080"}}
	client := NewDeviceFlowClientWithDeviceIDAndProxyURL(cfg, "device-1", proxyServer.URL)

	resp, errDo := client.httpClient.Get("http://example.com/test")
	if errDo != nil {
		t.Fatalf("client.Get returned error: %v", errDo)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
	if got := atomic.LoadInt32(&proxyCalls); got != 1 {
		t.Fatalf("proxy calls = %d, want 1", got)
	}
}
