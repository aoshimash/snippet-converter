package codesnippet

import "strings"

type ProgrammingLang struct {
	name         string
	fileSuffixes []string
	commentStart string
	commentEnd   string
}

var ProgrammingLangs = []*ProgrammingLang{
	{
		name:         "c",
		fileSuffixes: []string{".c"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "cpp",
		fileSuffixes: []string{".cc", ".cpp"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "java",
		fileSuffixes: []string{".java"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "go",
		fileSuffixes: []string{".go"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "scala",
		fileSuffixes: []string{".scala"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "groovy",
		fileSuffixes: []string{".groovy"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "php",
		fileSuffixes: []string{".php"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "javascript",
		fileSuffixes: []string{".js"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "css",
		fileSuffixes: []string{".css"},
		commentStart: "/*",
		commentEnd:   "*/",
	},
	{
		name:         "python",
		fileSuffixes: []string{".py"},
		commentStart: "'''",
		commentEnd:   "'''",
	},
	{
		name:         "bash",
		fileSuffixes: []string{".sh"},
		commentStart: "<<COMMENT",
		commentEnd:   "COMMENT",
	},
	{
		name:         "haskell",
		fileSuffixes: []string{".hs", ".lhs"},
		commentStart: "{-",
		commentEnd:   "-}",
	},
	{
		name:         "html",
		fileSuffixes: []string{".html"},
		commentStart: "<!--",
		commentEnd:   "-->",
	},
}

func getProgrammingLang(filePath string) (*ProgrammingLang, error) {
	for _, lang := range ProgrammingLangs {
		for _, suffix := range lang.fileSuffixes {
			if strings.HasSuffix(filePath, suffix) {
				return lang, nil
			}
		}
	}
	return nil, ErrUnsupportedFileExtension
}
