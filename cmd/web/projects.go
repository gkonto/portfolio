package main

type projectCardData struct {
	Title       string
	Label       string
	Summary     string
	Stack       string
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
		Stack:       "stack: Go / Bubble Tea / C++ include analysis",
		ImagePath:   "/static/img/inclens-topdown.gif",
		ImageAlt:    "IncLens top-down tree view demo",
		ProjectHref: "/projects/inclens",
	},
	{
		Title:       "Monkey in C++",
		Label:       "lang://monkey-interpreter",
		Summary:     "A modern C++ implementation of the Monkey language with a Pratt parser, tree-walking evaluator, REPL, tests, and runtime-focused optimization work.",
		Stack:       "stack: C++ / CMake / Catch2 / interpreter internals",
		ImagePath:   "/static/img/monkey-in-cpp-card.svg",
		ImageAlt:    "Monkey in C++ project preview",
		ProjectHref: "/projects/monkey-in-cpp",
	},
}
