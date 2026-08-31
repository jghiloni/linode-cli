package tools

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func GetRootDir() (rootDir string, err error) {
	rootDir = "."
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}

	if !filepath.IsAbs(rootDir) {
		cwd := ""
		cwd, err = os.Getwd()
		if err != nil {
			return
		}

		rootDir, err = filepath.Abs(filepath.Join(cwd, rootDir))
		if err != nil {
			return
		}
	}

	return
}

func GetConfig(rootDir string) (cfg *GenConfig, err error) {
	cfgFile, err := os.ReadFile(filepath.Join(rootDir, ".gen-config.json"))
	if err != nil {
		return
	}

	cfg = new(GenConfig)
	err = json.Unmarshal(cfgFile, cfg)
	return
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustValue[T any](val T, err error) T {
	Must(err)
	return val
}
