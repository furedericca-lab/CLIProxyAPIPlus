package management

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	coreauth "github.com/router-for-me/CLIProxyAPI/v7/sdk/cliproxy/auth"
)

func normalizeAntigravityPrimaryEntries(entries []gin.H) {
	if len(entries) == 0 {
		return
	}

	antigravityEntries := make([]gin.H, 0)
	nextOrder := 1
	for _, entry := range entries {
		if !isAntigravityEntry(entry) {
			continue
		}
		antigravityEntries = append(antigravityEntries, entry)
		if orderValue, ok := antigravityEntryOrder(entry); ok && orderValue >= nextOrder {
			nextOrder = orderValue + 1
		}
	}
	if len(antigravityEntries) == 0 {
		return
	}

	primaryEntry := selectAntigravityPrimaryEntry(antigravityEntries)
	for _, entry := range antigravityEntries {
		primaryInfo, _ := entry["primary_info"].(gin.H)
		if primaryInfo == nil {
			primaryInfo = gin.H{"order": nextOrder}
			entry["primary_info"] = primaryInfo
			nextOrder++
		} else if orderValue, ok := antigravityEntryOrder(entry); !ok || orderValue <= 0 {
			primaryInfo["order"] = nextOrder
			nextOrder++
		}
		primaryInfo["is_primary"] = primaryEntry != nil && antigravityEntryIdentity(entry) == antigravityEntryIdentity(primaryEntry)
	}
}

func selectAntigravityPrimaryEntry(entries []gin.H) gin.H {
	var primaryEntry gin.H
	for _, entry := range entries {
		if isDisabledAntigravityEntry(entry) {
			continue
		}
		if primaryEntry == nil {
			primaryEntry = entry
			continue
		}
		if compareAntigravityPrimaryEntry(entry, primaryEntry) < 0 {
			primaryEntry = entry
		}
	}
	return primaryEntry
}

func compareAntigravityPrimaryEntry(left, right gin.H) int {
	leftOrder, leftHasOrder := antigravityEntryOrder(left)
	rightOrder, rightHasOrder := antigravityEntryOrder(right)
	if leftHasOrder && rightHasOrder && leftOrder != rightOrder {
		return leftOrder - rightOrder
	}
	if leftHasOrder != rightHasOrder {
		if leftHasOrder {
			return -1
		}
		return 1
	}

	leftName := strings.ToLower(strings.TrimSpace(antigravityEntryString(left, "name")))
	rightName := strings.ToLower(strings.TrimSpace(antigravityEntryString(right, "name")))
	if leftName != rightName {
		if leftName < rightName {
			return -1
		}
		return 1
	}

	leftID := strings.ToLower(strings.TrimSpace(antigravityEntryIdentity(left)))
	rightID := strings.ToLower(strings.TrimSpace(antigravityEntryIdentity(right)))
	if leftID < rightID {
		return -1
	}
	if leftID > rightID {
		return 1
	}
	return 0
}

func antigravityEntryOrder(entry gin.H) (int, bool) {
	if entry == nil {
		return 0, false
	}
	primaryInfo, ok := entry["primary_info"].(gin.H)
	if !ok || primaryInfo == nil {
		return 0, false
	}
	switch value := primaryInfo["order"].(type) {
	case int:
		if value > 0 {
			return value, true
		}
	case int32:
		if value > 0 {
			return int(value), true
		}
	case int64:
		if value > 0 {
			return int(value), true
		}
	case float64:
		if value > 0 {
			return int(value), true
		}
	}
	return 0, false
}

func isAntigravityEntry(entry gin.H) bool {
	return strings.EqualFold(strings.TrimSpace(antigravityEntryString(entry, "type")), "antigravity") ||
		strings.EqualFold(strings.TrimSpace(antigravityEntryString(entry, "provider")), "antigravity")
}

func isDisabledAntigravityEntry(entry gin.H) bool {
	if entry == nil {
		return true
	}
	if disabled, ok := entry["disabled"].(bool); ok && disabled {
		return true
	}
	status := strings.TrimSpace(antigravityEntryString(entry, "status"))
	return strings.EqualFold(status, string(coreauth.StatusDisabled))
}

func antigravityEntryString(entry gin.H, key string) string {
	if entry == nil {
		return ""
	}
	if value, ok := entry[key].(string); ok {
		return value
	}
	return ""
}

func antigravityEntryIdentity(entry gin.H) string {
	if entry == nil {
		return ""
	}
	if id := strings.TrimSpace(antigravityEntryString(entry, "id")); id != "" {
		return id
	}
	if name := strings.TrimSpace(antigravityEntryString(entry, "name")); name != "" {
		return name
	}
	return antigravityEntryString(entry, "path")
}

func hasExplicitAntigravityPrimary(auths []*coreauth.Auth) bool {
	for _, auth := range auths {
		if auth == nil || !strings.EqualFold(strings.TrimSpace(auth.Provider), "antigravity") {
			continue
		}
		if auth.PrimaryInfo != nil && auth.PrimaryInfo.IsPrimary {
			return true
		}
	}
	return false
}

func ensureAntigravityPrimaryInfoEntry(entry gin.H, auth *coreauth.Auth, fallbackOrder int, primaryAlreadyAssigned bool) gin.H {
	if entry == nil || auth == nil {
		return entry
	}
	if !strings.EqualFold(strings.TrimSpace(auth.Provider), "antigravity") {
		return entry
	}
	if _, exists := entry["primary_info"]; exists {
		return entry
	}
	isPrimary := !primaryAlreadyAssigned && !auth.Disabled && auth.Status != coreauth.StatusDisabled
	entry["primary_info"] = gin.H{
		"is_primary": isPrimary,
		"order":      fallbackOrder,
	}
	return entry
}

// GetAuthFileModels returns the models supported by a specific auth file

func (h *Handler) ensureSoleAntigravityPrimary(ctx context.Context, primaryAuth *coreauth.Auth) {
	if h.authManager == nil || primaryAuth == nil {
		return
	}
	auths := h.authManager.List()
	for _, auth := range auths {
		if auth == nil || auth.ID == primaryAuth.ID {
			continue
		}
		if !strings.EqualFold(strings.TrimSpace(auth.Provider), "antigravity") {
			continue
		}
		shouldDemote := false
		if auth.PrimaryInfo != nil {
			shouldDemote = auth.PrimaryInfo.IsPrimary
		} else if !auth.Disabled && auth.Status != coreauth.StatusDisabled {
			shouldDemote = true
		}
		if shouldDemote {
			auth.Disabled = true
			auth.Status = coreauth.StatusDisabled
			auth.StatusMessage = "demoted via primary handoff"
			if auth.PrimaryInfo == nil {
				auth.PrimaryInfo = &coreauth.PrimaryInfo{}
			}
			auth.PrimaryInfo.IsPrimary = false
			auth.UpdatedAt = time.Now()
			_, _ = h.authManager.Update(ctx, auth)
		}
	}
	primaryAuth.Disabled = false
	primaryAuth.Status = coreauth.StatusActive
	primaryAuth.StatusMessage = ""
	primaryAuth.Unavailable = false
	if primaryAuth.PrimaryInfo != nil {
		primaryAuth.PrimaryInfo.IsPrimary = true
	}
	primaryAuth.UpdatedAt = time.Now()
	_, _ = h.authManager.Update(ctx, primaryAuth)
}

func (h *Handler) initAntigravityPrimaryInfo(ctx context.Context, record *coreauth.Auth) {
	if h == nil || h.cfg == nil {
		return
	}
	if record == nil || !strings.EqualFold(strings.TrimSpace(record.Provider), "antigravity") {
		return
	}
	existingPrimary := false
	maxOrder := 0
	if h.authManager != nil {
		for _, auth := range h.authManager.List() {
			if auth == nil || !strings.EqualFold(strings.TrimSpace(auth.Provider), "antigravity") {
				continue
			}
			if auth.ID == record.ID {
				continue
			}
			if auth.PrimaryInfo != nil {
				if auth.PrimaryInfo.Order > maxOrder {
					maxOrder = auth.PrimaryInfo.Order
				}
				if auth.PrimaryInfo.IsPrimary {
					existingPrimary = true
				}
				continue
			}
			if !auth.Disabled && auth.Status != coreauth.StatusDisabled {
				existingPrimary = true
			}
		}
	}
	if existingPrimary {
		record.PrimaryInfo = &coreauth.PrimaryInfo{
			IsPrimary: false,
			Order:     maxOrder + 1,
		}
		record.Disabled = true
		record.Status = coreauth.StatusDisabled
	} else {
		record.PrimaryInfo = &coreauth.PrimaryInfo{
			IsPrimary: true,
			Order:     maxOrder + 1,
		}
		record.Disabled = false
		record.Status = coreauth.StatusActive
	}
}
