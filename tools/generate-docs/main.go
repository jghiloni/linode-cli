package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jghiloni/linodectl/commands"
	"github.com/jghiloni/linodectl/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const indexFile = "index.md"

var relativeLinkFinder = regexp.MustCompile(`\]\(([\w_\-]+?\.md)\)`)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	rootDir := tools.MustValue(tools.GetRootDir())
	rootCmd := commands.Root()

	generateDocs(ctx, rootDir, rootCmd)
}

func generateDocs(ctx context.Context, rootDir string, cmd *cobra.Command) {
	checkCtx(ctx)
	docsDir := filepath.Join(rootDir, "docs", "commands")
	tools.Must(os.MkdirAll(docsDir, os.ModePerm))
	tools.Must(doc.GenMarkdownTreeCustom(cmd, docsDir, prepender, cobraLinkFunc))
	tools.Must(reorganizeFiles(docsDir, cmd.Name()))
}

func checkCtx(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}

	select {
	case <-ctx.Done():
		panic(ctx.Err())
	default:
		return
	}
}

func reorganizeFiles(docsDir, rootName string) error {
	if err := moveFile(docsDir, filepath.Join(docsDir, rootName+".md"), filepath.Join(docsDir, indexFile)); err != nil {
		return err
	}

	allDocsFiles, _ := filepath.Glob(fmt.Sprintf("%s/%s_*.md", docsDir, rootName))
	for _, file := range allDocsFiles {
		name := filepath.Base(file)
		parts := strings.Split(name, "_")
		if len(parts) < 2 {
			continue
		}
		var pathSegments []string
		if len(parts) == 2 {
			pathSegments = []string{strings.TrimSuffix(parts[1], ".md"), indexFile}
		} else {
			pathSegments = parts[1:]
		}

		newPath := filepath.Join(append([]string{docsDir}, pathSegments...)...)
		if err := moveFile(docsDir, file, newPath); err != nil {
			return err
		}
	}
	return nil
}

func moveFile(rootDir, oldFile, newFile string) error {
	fixed, err := fixAllLinks(rootDir, oldFile)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(newFile), os.ModePerm); err != nil {
		return err
	}

	if err = os.WriteFile(newFile, fixed, os.ModePerm^0o111); err != nil { // create the file with standard permissions with execute bits removed
		return err
	}

	return os.Remove(oldFile)
}

func fixAllLinks(rootDir, filename string) ([]byte, error) {
	relpath, err := filepath.Rel(rootDir, filename)
	if err != nil {
		return nil, err
	}
	relFile := getRelativePath(relpath)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	newData := relativeLinkFinder.ReplaceAllFunc(data, func(match []byte) []byte {
		linkPath := string(match[2 : len(match)-1]) // removes the leading '](' and trailing ')'
		relativeLink, _ := filepath.Rel(filepath.Dir(relFile), getRelativePath(linkPath))
		if relativeLink == "." {
			relativeLink = indexFile
		}
		if !strings.HasPrefix(relativeLink, ".") {
			relativeLink = "./" + relativeLink
		}

		return fmt.Appendf(make([]byte, 0, len(relativeLink)+3), "](%s)", relativeLink)
	})

	return newData, nil
}

func getRelativePath(fileName string) string {
	segments := strings.Split(fileName, "_")
	switch len(segments) {
	case 1:
		segments = []string{indexFile}
	case 2:
		segments = []string{strings.TrimSuffix(segments[1], ".md"), indexFile}
	default:
		segments = segments[1:]
	}

	return filepath.Join(segments...)
}

func prepender(s string) string {
	filename := filepath.Base(s)
	parts := strings.Split(filename, "_")
	if len(parts) == 2 {
		groupCmd := strings.TrimSuffix(parts[1], ".md")
		displayName, found := commands.DisplayNames.LookupLongName(groupCmd)
		if found {
			return fmt.Sprintf("# %s\n\n", displayName)
		}
	}
	return ""
}

func cobraLinkFunc(s string) string {
	return s
}
