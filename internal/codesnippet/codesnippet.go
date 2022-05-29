package codesnippet

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"strings"
)

var (
	headers = []map[string]string{
		{"start": "/*", "end": "*/"},           // c, cpp, java, go, scala, groovy, php, javascript, css
		{"start": "'''", "end": "'''"},         // python
		{"start": "<<HEADER", "end": "HEADER"}, // bash
		{"start": "{-", "end": "-}"},           // Haskell
		{"start": "<!--", "end": "-->"},        // html
		{"start": "###", "end": "###"},         // etc
		{"start": "///", "end": "///"},         // etc
	}
)

var (
	ErrInsufficientHeader = errors.New("Insufficient Header Information")
	ErrEmptyBody          = errors.New("Body is empty")
)

type codeSnippetValue struct {
	prefix, description string
	body                []string
}

// ファイルの情報を格納する構造体
type CodeSnippet struct {
	key   string
	value codeSnippetValue
}

// ファイルの情報を取得
func NewCodeSnippet(path string) (CodeSnippet, error) {
	snippet := CodeSnippet{}

	data, err := os.Open(path)
	defer data.Close()

	if err != nil {
		return CodeSnippet{}, err
	}

	header_start_flag := false
	header_end_flag := false

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		text := scanner.Text()
		if !header_start_flag {
			for _, headerComment := range headers {
				if strings.HasPrefix(text, headerComment["start"]) {
					header_start_flag = true
					break
				}
			}
		} else if !header_end_flag {
			// headerの終端であるか確認
			for _, headerComment := range headers {
				if strings.HasPrefix(text, headerComment["end"]) {
					if snippet.key == "" || snippet.value.prefix == "" || snippet.value.description == "" {
						// headerの終端まできてるのに、key, prefix, descriptionのいずれかが空の場合は空のCodeSnippetをリターンする
						return CodeSnippet{}, ErrInsufficientHeader
					}
					header_end_flag = true
					break
				}
			}
			//
			if strings.HasPrefix(text, "key:") {
				snippet.key = strings.TrimPrefix(text, "key: ")
			} else if strings.HasPrefix(text, "prefix:") {
				snippet.value.prefix = strings.TrimPrefix(text, "prefix: ")
			} else if strings.HasPrefix(text, "description:") {
				snippet.value.description = strings.TrimPrefix(text, "description: ")
			}
		} else {
			snippet.value.body = append(snippet.value.body, text)
		}
	}

	// bodyが空の場合は空CodeSnippetをリターン
	if len(snippet.value.body) == 0 {
		return CodeSnippet{}, ErrEmptyBody
	}

	return snippet, err
}

func NewCodeSnippets(filePathes []string) []CodeSnippet {
	snippets := []CodeSnippet{}

	for _, filePath := range filePathes {
		snippet, err := NewCodeSnippet(filePath)
		if errors.Is(err, fs.ErrNotExist) ||
			errors.Is(err, ErrEmptyBody) ||
			errors.Is(err, ErrInsufficientHeader) {
			continue
		}
		snippets = append(snippets, snippet)
	}

	return snippets
}

func GetVSCodeSnippetsJSON(snippets []CodeSnippet) ([]byte, error) {
	jsonMap := map[string]interface{}{}

	for _, snippet := range snippets {
		jsonMap[snippet.key] = map[string]interface{}{
			"prefix":      snippet.value.prefix,
			"body":        snippet.value.body,
			"description": snippet.value.description,
		}
	}

	return json.Marshal(jsonMap)
}
