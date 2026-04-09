package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"portfolio.gkontogiannis.net/ui"
)

type exportPage struct {
	route    string
	template string
	data     any
}

func (app *application) exportSite(outputDir string) error {
	if err := os.RemoveAll(outputDir); err != nil {
		return err
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}

	pages := []exportPage{
		{
			route:    "/",
			template: "home.tmpl.html",
			data: homePageData{
				Projects: portfolioProjects,
			},
		},
		{
			route:    "/contact",
			template: "contact.tmpl.html",
			data:     nil,
		},
		{
			route:    "/projects/inclens",
			template: "inclens.tmpl.html",
			data:     inclensPage,
		},
		{
			route:    "/projects/monkey-in-cpp",
			template: "monkey-in-cpp.tmpl.html",
			data:     monkeyInCPPPage,
		},
	}

	for _, page := range pages {
		content, err := app.renderPage(page.template, page.data)
		if err != nil {
			return err
		}

		targetPath := exportTargetPath(outputDir, page.route)
		if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
			return err
		}

		if err := os.WriteFile(targetPath, content, 0o644); err != nil {
			return err
		}
	}

	return copyEmbeddedDir(ui.Files, "static", filepath.Join(outputDir, "static"))
}

func exportTargetPath(outputDir, route string) string {
	if route == "/" {
		return filepath.Join(outputDir, "index.html")
	}

	cleanRoute := strings.TrimPrefix(route, "/")
	return filepath.Join(outputDir, filepath.FromSlash(cleanRoute), "index.html")
}

func copyEmbeddedDir(embeddedFS fs.FS, sourceDir, targetDir string) error {
	return fs.WalkDir(embeddedFS, sourceDir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(targetDir, relativePath)
		if entry.IsDir() {
			return os.MkdirAll(targetPath, 0o755)
		}

		content, err := fs.ReadFile(embeddedFS, path)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath, content, 0o644)
	})
}
