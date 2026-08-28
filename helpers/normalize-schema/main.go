package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/forPelevin/gomoji"
)

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
					tags := make([]string, 0, len(kvs))
					for _, maybeTag := range kvs {
						if tag, ok := maybeTag.(string); ok {
							tags = append(tags, strings.ToLower(tag))
						}
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
	var parsed any
	if err := json.NewDecoder(os.Stdin).Decode(&parsed); err != nil {
		log.Fatal(err)
	}

	scrubbed := stripAny(parsed)
	if err := json.NewEncoder(os.Stdout).Encode(scrubbed); err != nil {
		log.Fatal(err)
	}
}
