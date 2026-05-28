package iflow

import (
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v7/internal/registry"
	"github.com/router-for-me/CLIProxyAPI/v7/internal/thinking"
	"github.com/tidwall/gjson"
)

func TestApplyKimiFallbackThinking(t *testing.T) {
	applier := NewApplier()
	modelInfo := &registry.ModelInfo{
		ID:       "kimi-k2.5",
		Thinking: &registry.ThinkingSupport{},
	}

	tests := []struct {
		name       string
		config     thinking.ThinkingConfig
		wantEnable bool
	}{
		{
			name:       "none disables thinking",
			config:     thinking.ThinkingConfig{Mode: thinking.ModeNone},
			wantEnable: false,
		},
		{
			name:       "auto enables thinking",
			config:     thinking.ThinkingConfig{Mode: thinking.ModeAuto},
			wantEnable: true,
		},
		{
			name:       "level enables thinking",
			config:     thinking.ThinkingConfig{Mode: thinking.ModeLevel, Level: thinking.LevelHigh},
			wantEnable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applier.Apply([]byte(`{"model":"kimi-k2.5"}`), tt.config, modelInfo)
			if err != nil {
				t.Fatalf("Apply() error = %v", err)
			}

			enableResult := gjson.GetBytes(got, "chat_template_kwargs.enable_thinking")
			if !enableResult.Exists() {
				t.Fatalf("enable_thinking missing in %s", string(got))
			}
			if enableResult.Bool() != tt.wantEnable {
				t.Fatalf("enable_thinking = %v, want %v", enableResult.Bool(), tt.wantEnable)
			}
			if gjson.GetBytes(got, "chat_template_kwargs.clear_thinking").Exists() {
				t.Fatalf("clear_thinking should not be set for Kimi fallback: %s", string(got))
			}
		})
	}
}

func TestApplyKimiWithoutThinkingSupportIsNoop(t *testing.T) {
	applier := NewApplier()
	modelInfo := &registry.ModelInfo{ID: "kimi-k2.5"}
	body := []byte(`{"model":"kimi-k2.5"}`)

	got, err := applier.Apply(body, thinking.ThinkingConfig{Mode: thinking.ModeAuto}, modelInfo)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	if string(got) != string(body) {
		t.Fatalf("Apply() = %s, want unchanged %s", string(got), string(body))
	}
}
