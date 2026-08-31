package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type GenConfig struct {
	Name           string `json:"name"`
	Module         string `json:"module"`
	SpecRepository string `json:"api_repository"`
	SpecVersion    string `json:"api_version"`
}

func ParseGenConfig(ctx context.Context, reader io.Reader) (*GenConfig, error) {
	var cfg GenConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	return &cfg, err
}

func GetOpenAPISpec(ctx context.Context, config *GenConfig, targetDir string) (string, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/refs/tags/%s/openapi.json", config.SpecRepository, config.SpecVersion)

	fp, err := os.Create(filepath.Join(targetDir, "openapi.json"))
	if err != nil {
		return "", err
	}
	defer fp.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", resp.Status)
	}

	_, err = io.Copy(fp, resp.Body)
	return fp.Name(), err
}

func NormalizeTags(file string) (string, error) {
	var rawSpec any

	fp, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	if err = json.NewDecoder(fp).Decode(&rawSpec); err != nil {
		return "", err
	}

	newSpecFile, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	defer newSpecFile.Close()

	err = json.NewEncoder(newSpecFile).Encode(normalizeTags(rawSpec))
	return newSpecFile.Name(), err
}

func normalizeTags(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		newMap := make(map[string]any, len(typed))
		for key, val := range typed {
			if key == "tags" {
				if maybeTags, ok := val.([]any); ok {
					tags := make([]string, 0, len(maybeTags))
					for _, maybeTag := range maybeTags {
						if tag, ok := maybeTag.(string); ok {
							tags = append(tags, strings.ToLower(tag))
						}
					}
					newMap[key] = tags
					continue
				}
			}
			newMap[key] = normalizeTags(val)
		}
		return newMap
	case []any:
		newSlice := make([]any, len(typed))
		for i, v := range typed {
			newSlice[i] = normalizeTags(v)
		}
		return newSlice
	}

	return value
}
