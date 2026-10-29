package config

import "sync"

var autoUpdateMu sync.RWMutex

// IsAutoUpdateEnabled reports whether installed apps are updated automatically. Safe to call from
// any goroutine (the hourly job reads it while the settings endpoint may be changing it).
func IsAutoUpdateEnabled() bool {
	autoUpdateMu.RLock()
	defer autoUpdateMu.RUnlock()
	return AppInfo.AutoUpdateEnabled
}

// SetAutoUpdateEnabled changes the flag in memory only; call SaveSetup to persist it.
func SetAutoUpdateEnabled(enabled bool) {
	autoUpdateMu.Lock()
	defer autoUpdateMu.Unlock()
	AppInfo.AutoUpdateEnabled = enabled
}
