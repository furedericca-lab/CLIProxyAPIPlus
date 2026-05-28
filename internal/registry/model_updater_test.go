package registry

import (
	"strings"
	"testing"
)

func TestModelsURLsUseForkSource(t *testing.T) {
	if len(modelsURLs) == 0 {
		t.Fatal("modelsURLs is empty")
	}
	for _, url := range modelsURLs {
		if !strings.Contains(url, "furedericca-lab/models") {
			t.Fatalf("model source %q does not use the maintained fork", url)
		}
		if strings.Contains(url, "router-for-me/models") || strings.Contains(url, "models.router-for.me") {
			t.Fatalf("model source %q points at router catalog", url)
		}
	}
}
