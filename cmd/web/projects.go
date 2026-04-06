package main

type projectCardData struct {
	Title       string
	Label       string
	Summary     string
	Tags        []string
	ImagePath   string
	ImageAlt    string
	ProjectHref string
}

type homePageData struct {
	Projects []projectCardData
}

var portfolioProjects = []projectCardData{
	{
		Title:       "IncLens",
		Label:       "tool://inclens",
		Summary:     "A terminal UI for exploring C++ include relationships from a preprocessed .ii file, with both tree and flamegraph views.",
		Tags:        []string{"Go", "C++", "TUI"},
		ImagePath:   "/static/img/inclens-topdown.gif",
		ImageAlt:    "IncLens top-down tree view demo",
		ProjectHref: "/projects/inclens",
	},
	{
		Title:       "Monkey in C++",
		Label:       "lang://monkey-interpreter",
		Summary:     "A modern C++ implementation of the Monkey language with a Pratt parser, tree-walking evaluator, REPL, tests, and runtime-focused optimization work.",
		Tags:        []string{"C++", "CMake", "Catch2", "Interpreters"},
		ImagePath:   "/static/img/monkey-in-cpp-card.png",
		ImageAlt:    "Illustration showing 2 plus 3 times 4 and its syntax tree",
		ProjectHref: "/projects/monkey-in-cpp",
	},
}
