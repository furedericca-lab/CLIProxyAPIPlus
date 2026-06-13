package config

import "strings"

type OAuthEndpointConfig struct {
	ApiBaseURL         string `yaml:"api-base-url,omitempty" json:"api-base-url,omitempty"`
	AuthorizeURL       string `yaml:"authorize-url,omitempty" json:"authorize-url,omitempty"`
	TokenURL           string `yaml:"token-url,omitempty" json:"token-url,omitempty"`
	RefreshURL         string `yaml:"refresh-url,omitempty" json:"refresh-url,omitempty"`
	UserinfoURL        string `yaml:"userinfo-url,omitempty" json:"userinfo-url,omitempty"`
	DeviceAuthorizeURL string `yaml:"device-authorize-url,omitempty" json:"device-authorize-url,omitempty"`
}

func (c *OAuthEndpointConfig) ApplyDefaults(defaults OAuthEndpointConfig) OAuthEndpointConfig {
	result := *c
	if result.ApiBaseURL == "" {
		result.ApiBaseURL = defaults.ApiBaseURL
	}
	if result.AuthorizeURL == "" {
		result.AuthorizeURL = defaults.AuthorizeURL
	}
	if result.TokenURL == "" {
		result.TokenURL = defaults.TokenURL
	}
	if result.RefreshURL == "" {
		result.RefreshURL = defaults.RefreshURL
	}
	if result.UserinfoURL == "" {
		result.UserinfoURL = defaults.UserinfoURL
	}
	if result.DeviceAuthorizeURL == "" {
		result.DeviceAuthorizeURL = defaults.DeviceAuthorizeURL
	}
	return result
}

// SanitizeOAuthModelAlias normalizes and deduplicates global OAuth model name aliases.
// It trims whitespace, normalizes channel keys to lower-case, drops empty entries,
// allows multiple source models to share the same alias, and ensures each name+alias
// combination is unique within each channel.
// It also injects default aliases for channels that have built-in defaults (e.g. Kiro)
// when no user-configured aliases exist for those channels.
func (cfg *Config) SanitizeOAuthModelAlias() {
	if cfg == nil {
		return
	}

	if cfg.OAuthModelAlias == nil {
		cfg.OAuthModelAlias = make(map[string][]OAuthModelAlias)
	}
	hasChannel := func(channel string) bool {
		for k := range cfg.OAuthModelAlias {
			if strings.EqualFold(strings.TrimSpace(k), channel) {
				return true
			}
		}
		return false
	}
	if !hasChannel("kiro") {
		cfg.OAuthModelAlias["kiro"] = defaultKiroAliases()
	}
	if !hasChannel("github-copilot") {
		cfg.OAuthModelAlias["github-copilot"] = defaultGitHubCopilotAliases()
	}

	if len(cfg.OAuthModelAlias) == 0 {
		return
	}
	out := make(map[string][]OAuthModelAlias, len(cfg.OAuthModelAlias))
	for rawChannel, aliases := range cfg.OAuthModelAlias {
		channel := strings.ToLower(strings.TrimSpace(rawChannel))
		if channel == "" {
			continue
		}
		// Preserve explicit empty markers so default injection will not re-add them.
		if len(aliases) == 0 {
			out[channel] = nil
			continue
		}
		seenNameAlias := make(map[string]struct{}, len(aliases))
		clean := make([]OAuthModelAlias, 0, len(aliases))
		for _, entry := range aliases {
			name := strings.TrimSpace(entry.Name)
			alias := strings.TrimSpace(entry.Alias)
			if name == "" || alias == "" {
				continue
			}
			if strings.EqualFold(name, alias) {
				continue
			}
			nameAliasKey := strings.ToLower(name + "::" + alias)
			if _, ok := seenNameAlias[nameAliasKey]; ok {
				continue
			}
			seenNameAlias[nameAliasKey] = struct{}{}
			clean = append(clean, OAuthModelAlias{Name: name, Alias: alias, Fork: entry.Fork})
		}
		if len(clean) > 0 {
			out[channel] = clean
		}
	}
	cfg.OAuthModelAlias = out
}

func (cfg *Config) NormalizeOAuthEndpointOverrides() {
	if cfg == nil || len(cfg.OAuthEndpointOverrides) == 0 {
		return
	}
	normalized := make(map[string]OAuthEndpointConfig, len(cfg.OAuthEndpointOverrides))
	for provider, ep := range cfg.OAuthEndpointOverrides {
		normalizedProvider := strings.ToLower(strings.TrimSpace(provider))
		if normalizedProvider == "" {
			continue
		}
		ep.ApiBaseURL = strings.TrimSpace(ep.ApiBaseURL)
		ep.AuthorizeURL = strings.TrimSpace(ep.AuthorizeURL)
		ep.TokenURL = strings.TrimSpace(ep.TokenURL)
		ep.RefreshURL = strings.TrimSpace(ep.RefreshURL)
		ep.UserinfoURL = strings.TrimSpace(ep.UserinfoURL)
		ep.DeviceAuthorizeURL = strings.TrimSpace(ep.DeviceAuthorizeURL)
		normalized[normalizedProvider] = ep
	}
	cfg.OAuthEndpointOverrides = normalized
}

func (cfg *Config) GetOAuthEndpointOverride(provider string) OAuthEndpointConfig {
	if cfg == nil {
		return OAuthEndpointConfig{}
	}
	normalizedProvider := strings.ToLower(strings.TrimSpace(provider))
	if cfg.OAuthEndpointOverrides != nil {
		if ep, ok := cfg.OAuthEndpointOverrides[normalizedProvider]; ok {
			return ep
		}
	}
	return OAuthEndpointConfig{}
}

// NormalizeOAuthExcludedModels cleans provider -> excluded models mappings by normalizing provider keys
// and applying model exclusion normalization to each entry.
func NormalizeOAuthExcludedModels(entries map[string][]string) map[string][]string {
	if len(entries) == 0 {
		return nil
	}
	out := make(map[string][]string, len(entries))
	for provider, models := range entries {
		key := strings.ToLower(strings.TrimSpace(provider))
		if key == "" {
			continue
		}
		normalized := NormalizeExcludedModels(models)
		if len(normalized) == 0 {
			continue
		}
		out[key] = normalized
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
