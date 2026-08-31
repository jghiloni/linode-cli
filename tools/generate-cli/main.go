package main

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/forPelevin/gomoji"
	"github.com/jghiloni/linodectl/tools"
)

var replacedFiles = []string{
	// "commands/auth.go",
	// "runtime/config.go",
}

func stripAny(val any) any {
	switch v := val.(type) {
	case string:
		return gomoji.ReplaceEmojisWithSlug(v)
	case []any:
		newV := make([]any, len(v))
		for i := range v {
			newV[i] = stripAny(v[i])
		}
		return newV
	case map[string]any:
		newV := make(map[string]any, len(v))
		for k, kv := range v {
			if k == "tags" {
				if kvs, ok := kv.([]any); ok {
					tags := make([]any, 0, len(kvs))
					for _, maybeTag := range kvs {
						if tag, ok := maybeTag.(string); ok {
							tags = append(tags, strings.ToLower(tag))
							continue
						}
						tags = append(tags, stripAny(maybeTag))
					}
					newV[k] = tags
					continue
				}
			}
			newV[stripAny(k).(string)] = stripAny(kv)
		}
		return newV
	}

	return val
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	rootDir := tools.MustValue(tools.GetRootDir())
	cfg := tools.MustValue(tools.GetConfig(rootDir))

	specPath := tools.MustValue(tools.GetOpenAPISpec(context.Background(), cfg, rootDir))
	spec := tools.MustValue(os.ReadFile(specPath))

	var parsed any
	tools.Must(json.Unmarshal(spec, &parsed))
	scrubbed := stripAny(parsed)

	encoded := tools.MustValue(json.Marshal(scrubbed))
	tools.Must(os.WriteFile(specPath, encoded, 0600))

	toolArgs := []string{"tool", "onlycli", "generate", "--name", cfg.Name, "--module", cfg.Module, "--out", rootDir, "--spec", specPath}
	cmd := exec.CommandContext(ctx, "go", toolArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	tools.Must(cmd.Run())

	for _, glob := range replacedFiles {
		files := tools.MustValue(filepath.Glob(filepath.Join(rootDir, glob)))
		for _, file := range files {
			tools.Must(os.Remove(file))
		}
	}

	tidyCmd := exec.CommandContext(ctx, "go", "mod", "tidy")
	tools.Must(tidyCmd.Run())
}
