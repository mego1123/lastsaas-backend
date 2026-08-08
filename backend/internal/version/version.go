package version

import (
        "log/slog"
        "os"
        "path/filepath"
        "strings"
)

// Current holds the version loaded at startup.
var Current string

// buildVersion is set at compile time via:
//
//      go build -ldflags "-X lastsaas/internal/version.buildVersion=1.00"
//
// This is the preferred mechanism — it bakes the version into the binary
// so downstream projects don't need to copy a VERSION file at runtime.
var buildVersion string

// Load returns the version string. Priority:
//  1. Compile-time ldflags (buildVersion)
//  2. VERSION file on disk (walking up from cwd)
//  3. "unknown"
func Load() string {
        if buildVersion != "" {
                Current = buildVersion
                return Current
        }

        dir, err := os.Getwd()
        if err != nil {
                slog.Warn("version: failed to get cwd, will not search for VERSION file", "error", err)
        }
        for i := 0; i < 5; i++ {
                data, err := os.ReadFile(filepath.Join(dir, "VERSION"))
                if err == nil {
                        Current = strings.TrimSpace(string(data))
                        return Current
                }
                parent := filepath.Dir(dir)
                if parent == dir {
                        break
                }
                dir = parent
        }
        Current = "unknown"
        return Current
}
